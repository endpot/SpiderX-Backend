package torrent

import (
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"github.com/anacrolix/torrent/bencode"
	"github.com/endpot/SpiderX-Backend/app/domain/model"
	"github.com/endpot/SpiderX-Backend/app/infra/db"
	customError "github.com/endpot/SpiderX-Backend/app/infra/util/error"
	"github.com/endpot/SpiderX-Backend/app/infra/util/fs"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"mime/multipart"
	"path"
)

const (
	TORRENT_PENDING_STATE = iota
	TORRENT_NORMAL_STATE
	TORRENT_DEAD_STATE
)

type BencodeTorrent struct {
	Announce  string `bencode:"announce"`
	CreatedBy string `bencode:"created by,omitempty"`
	CreatedAt int    `bencode:"creation date,omitempty"`
	Info      struct {
		Files []struct {
			Length int      `bencode:"length"`
			Path   []string `bencode:"path"`
		} `bencode:"files"`
		Name        string `bencode:"name"`
		Pieces      string `bencode:"pieces"`
		PieceLength int    `bencode:"piece length"`
		Private     int    `bencode:"private"`
		Source      string `bencode:"source"`
	} `bencode:"info"`
}

// 查询种子列表
func GetTorrentList(ctx *gin.Context) (model.TorrentSlice, error) {
	torrentSlice, err := model.Torrents().All(ctx, db.DB)

	return torrentSlice, err
}

// 预上传种子
func PreUploadTorrent(ctx *gin.Context) (*model.Torrent, *customError.CustomError) {
	// 检查携带种子文件
	torrentFileHeader, err := ctx.FormFile("torrent")
	if err != nil {
		return nil, customError.NewBadRequestError("TORRENT__EMPTY_TORRENT_FILE")
	}

	// 计算种子的 HASH 值
	infoHash, err := calTorrentHash(torrentFileHeader)
	if err != nil {
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
		if torrent.State == TORRENT_PENDING_STATE {
			return torrent, nil
		} else {
			return nil, customError.NewBadRequestError("TORRENT__TORRENT_EXISTS")
		}
	}

	// 种子不存在，插入数据库表
	torrent = &model.Torrent{
		InfoHash: infoHash,
	}
	if torrent.Insert(ctx, db.DB, boil.Infer()) != nil {
		return nil, customError.NewInternalServerError("TORRENT__INSERT_FAILED")
	}

	// 保存种子文件
	if saveTorrentFile(ctx, torrentFileHeader, torrent) != nil {
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

// 计算种子文件的 HASH 值
func calTorrentHash(torrent *multipart.FileHeader) (string, error) {
	// open file
	fileReader, err := torrent.Open()
	if err != nil {
		return "", err
	}

	// Decode
	// See https://godoc.org/github.com/anacrolix/torrent/bencode
	decoder := bencode.NewDecoder(fileReader)
	bencodeTorrent := &BencodeTorrent{}
	decodeErr := decoder.Decode(bencodeTorrent)
	if decodeErr != nil {
		return "", decodeErr
	}

	// marshal info part and calculate SHA1
	marshaledInfo, marshalErr := bencode.Marshal(bencodeTorrent.Info)
	if marshalErr != nil {
		return "", nil
	}

	return fmt.Sprintf("%x", sha1.Sum(marshaledInfo)), nil
}

// 保存种子到本地缓存目录和 OSS
func saveTorrentFile(ctx *gin.Context, torrentFile *multipart.FileHeader, torrent *model.Torrent) error {
	if torrent.ID <= 0 {
		return errors.New("invalid torrent")
	}

	torrentName := fmt.Sprintf("%v.torrent", torrent.ID)
	localPath := path.Join("./storage/cache/torrents/", torrentName)
	remotePath := "torrents/" + torrentName

	// 保存到本地
	if ctx.SaveUploadedFile(torrentFile, localPath) != nil {
		return errors.New("save to local failed")
	}

	// 保存到远端
	if fs.UploadToRemote(localPath, remotePath) != nil {
		return errors.New("save to remote failed")
	}

	return nil
}
