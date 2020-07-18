package torrent

import (
	customError "github.com/endpot/SpiderX-Backend/app/infra/util/error"
	"github.com/gin-gonic/gin"
	"strings"
)

type PreUploadTorrentRequest struct {
}

func (request *PreUploadTorrentRequest) ValidateRequest(ctx *gin.Context) *customError.CustomError {
	torrentFileHeader, err := ctx.FormFile("torrent")
	if err != nil {
		return customError.NewBadRequestError("TORRENT__EMPTY_TORRENT_FILE")
	}

	if !strings.HasSuffix(torrentFileHeader.Filename, ".torrent") {
		return customError.NewBadRequestError("TORRENT__INVALID_EXTENSION")
	}

	// TODO: 通过配置限制体积大小
	if torrentFileHeader.Size > 1024*1024*2 {
		return customError.NewBadRequestError("TORRENT__FILE_TOO_LARGE")
	}

	return nil
}

func (request *PreUploadTorrentRequest) ValidateAuth(ctx *gin.Context) *customError.CustomError {
	return nil
}
