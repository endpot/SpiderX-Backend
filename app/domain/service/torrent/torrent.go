package torrent

import (
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"github.com/anacrolix/torrent/bencode"
	torrentReq "github.com/endpot/SpiderX-Backend/app/controller/request/torrent"
	"github.com/endpot/SpiderX-Backend/app/domain/model"
	torrentExt "github.com/endpot/SpiderX-Backend/app/domain/modext/torrent"
	"github.com/endpot/SpiderX-Backend/app/infra/db"
	"github.com/endpot/SpiderX-Backend/app/infra/util/convert"
	customError "github.com/endpot/SpiderX-Backend/app/infra/util/error"
	"github.com/endpot/SpiderX-Backend/app/infra/util/fs"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"mime/multipart"
	"os"
	"path"
	"time"
)

type BencodeTorrent struct {
	Announce  string `bencode:"announce"`
	CreatedBy string `bencode:"created by,omitempty"`
	CreatedAt int    `bencode:"creation date,omitempty"`
	Info      struct {
		Files []struct {
			Length uint64   `bencode:"length"`
			Path   []string `bencode:"path"`
		} `bencode:"files"`
		Name        string `bencode:"name"`
		Pieces      string `bencode:"pieces"`
		PieceLength uint64 `bencode:"piece length"`
		Private     int    `bencode:"private"`
		Source      string `bencode:"source"`
	} `bencode:"info"`
}

// 查询种子列表
func GetTorrentList(ctx *gin.Context) (model.TorrentSlice, error) {
	torrentSlice, err := model.Torrents().All(ctx, db.DB)

	return torrentSlice, err
}

// 创建种子
func CreateTorrent(ctx *gin.Context, req *torrentReq.CreateTorrentRequest) (*model.Torrent, *customError.CustomError) {
	torrent, err := getTorrentByInfoHash(ctx, req.InfoHash)
	if err != nil || torrent == nil {
		return nil, customError.NewBadRequestError("TORRENT__INVALID_HASH")
	}

	torrent.CategoryID = req.CategoryID
	torrent.Title = req.Title
	torrent.SimpleDesc = req.SimpleDesc
	torrent.Description = req.Description
	torrent.IsAnonymous = convert.ParseBoolToUint8(req.IsAnonymous)
	torrent.PositionLevel = req.PositionLevel
	torrent.State = torrentExt.NORMAL_STATE
	// TODO: UploaderID/OwnerID 从 ctx 的用户信息中取
	torrent.UploaderID = 1
	torrent.OwnerID = 1
	torrent.CreatedAt = null.TimeFrom(time.Now().In(boil.GetLocation()))

	if _, err := torrent.Update(ctx, db.DB, boil.Infer()); err != nil {
		return nil, customError.NewBadRequestError("TORRENT__CREATE_FAILED")
	}

	return torrent, nil
}

// 预上传种子
func PreUploadTorrent(ctx *gin.Context, req *torrentReq.PreUploadTorrentRequest) (*model.Torrent, *customError.CustomError) {
	// 计算种子的 HASH 值
	bencodeTorrent, infoHash, err := repackTorrent(req.Torrent)
	if err != nil || bencodeTorrent == nil {
		return nil, customError.NewBadRequestError("TORRENT__HASH_FAILED")
	}

	// 检查种子是否已存在
	torrent, err := getTorrentByInfoHash(ctx, infoHash)
	if err != nil {
		// 查询数据库异常，此处认为种子已存在
		return nil, customError.NewBadRequestError("TORRENT__TORRENT_EXISTS")
	}

	// 种子已存在
	if torrent != nil {
		if torrent.State == torrentExt.PENDING_STATE {
			return torrent, nil
		} else {
			return nil, customError.NewBadRequestError("TORRENT__TORRENT_EXISTS")
		}
	}

	// 计算种子体积
	torrentSize, err := calTorrentSize(bencodeTorrent)
	if err != nil {
		return nil, customError.NewBadRequestError("TORRENT__HASH_FAILED")
	}

	// 种子不存在，插入数据库表
	torrent = &model.Torrent{
		InfoHash: infoHash,
		Size:     torrentSize,
	}
	if torrent.Insert(ctx, db.DB, boil.Infer()) != nil {
		return nil, customError.NewInternalServerError("TORRENT__INSERT_FAILED")
	}

	// 保存种子文件
	if saveTorrentFile(bencodeTorrent, torrent) != nil {
		// 保存失败的种子，从 DB 中硬删除
		_, _ = torrent.Delete(ctx, db.DB, true)
		return nil, customError.NewInternalServerError("TORRENT__SAVE_FAILED")
	}

	return torrent, nil
}

// 根据 HASH 值查找种子（忽略 ErrNoRows 错误）
func getTorrentByInfoHash(ctx *gin.Context, infoHash string) (*model.Torrent, error) {
	torrent, err := model.Torrents(
		model.TorrentWhere.InfoHash.EQ(infoHash),
	).One(ctx, db.DB)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return torrent, nil
}

// 改装种子文件，去除原 Tracker 信息，修改 Source 信息
func repackTorrent(fh *multipart.FileHeader) (*BencodeTorrent, string, error) {
	// open file
	fileReader, err := fh.Open()
	if err != nil {
		return nil, "", err
	}

	// Decode
	// See https://godoc.org/github.com/anacrolix/torrent/bencode
	decoder := bencode.NewDecoder(fileReader)
	bencodeTorrent := &BencodeTorrent{}
	decodeErr := decoder.Decode(bencodeTorrent)
	if decodeErr != nil {
		return nil, "", decodeErr
	}

	// Re-pack torrent
	// TODO: 根据配置修改
	bencodeTorrent.Announce = ""
	bencodeTorrent.Info.Source = "[Alpha] SpiderX"
	bencodeTorrent.Info.Private = 1

	// marshal info part and calculate SHA1
	marshaledInfo, marshalErr := bencode.Marshal(bencodeTorrent.Info)
	if marshalErr != nil {
		return nil, "", nil
	}

	return bencodeTorrent, fmt.Sprintf("%x", sha1.Sum(marshaledInfo)), nil
}

// 计算种子体积
func calTorrentSize(bencodeTorrent *BencodeTorrent) (uint64, error) {
	var totalSize uint64
	for _, file := range bencodeTorrent.Info.Files {
		totalSize += file.Length
	}

	return totalSize, nil
}

// 保存种子到本地缓存目录和 OSS
func saveTorrentFile(bencodeTorrent *BencodeTorrent, torrent *model.Torrent) error {
	if torrent.ID <= 0 {
		return errors.New("invalid torrent")
	}

	torrentName := fmt.Sprintf("%v.torrent", torrent.ID)
	localPath := path.Join("./storage/cache/torrents/", torrentName)
	remotePath := "torrents/" + torrentName

	// 保存到本地
	fileWriter, err := os.Create(localPath)
	if err != nil {
		return errors.New("create file failed")
	}

	encoder := bencode.NewEncoder(fileWriter)
	if err := encoder.Encode(bencodeTorrent); err != nil {
		return errors.New("save to local failed")
	}

	// 保存到远端
	if fs.UploadToRemote(localPath, remotePath) != nil {
		return errors.New("save to remote failed")
	}

	return nil
}
