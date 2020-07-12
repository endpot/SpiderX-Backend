package torrent

import (
	"crypto/sha1"
	"fmt"
	"github.com/anacrolix/torrent/bencode"
	"github.com/endpot/SpiderX-Backend/app/domain/model"
	"github.com/endpot/SpiderX-Backend/app/infra/db"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
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

func GetTorrentList(ctx *gin.Context) (model.TorrentSlice, error) {
	torrentSlice, err := model.Torrents().All(ctx, db.DB)

	return torrentSlice, err
}

func PreUploadTorrent(ctx *gin.Context) error {
	// 检查种子是否存在
	torrent, err := ctx.FormFile("torrent")
	if err != nil {
		return err
	}

	// 计算种子的 HASH 值
	infoHash, err := calTorrentHash(torrent)
	if err != nil {
		return err
	}
	log.Println("SHA1:", infoHash)

	// TODO: 检查种子是否已存在

	// TODO: 插入数据库表

	// TODO: 保存种子文件到缓存目录

	// TODO: 上传种子文件到远端

	// TODO: 返回操作结果
	return nil
}

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
