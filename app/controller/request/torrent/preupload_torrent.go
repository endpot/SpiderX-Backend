package torrent

import (
	customError "github.com/endpot/SpiderX-Backend/app/infra/util/error"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strings"
)

type PreUploadTorrentRequest struct {
	Torrent *multipart.FileHeader `form:"torrent" binding:"required"`
}

func (request *PreUploadTorrentRequest) ValidateRequest(ctx *gin.Context) *customError.CustomError {
	// 校验参数格式
	if err := ctx.ShouldBind(request); err != nil {
		return customError.NewBadRequestError("TORRENT__EMPTY_TORRENT_FILE")
	}

	if !strings.HasSuffix(request.Torrent.Filename, ".torrent") {
		return customError.NewBadRequestError("TORRENT__INVALID_EXTENSION")
	}

	// TODO: 通过配置限制体积大小
	if request.Torrent.Size > 1024*1024*2 {
		return customError.NewBadRequestError("TORRENT__FILE_TOO_LARGE")
	}

	return nil
}

func (request *PreUploadTorrentRequest) ValidateAuth(ctx *gin.Context) *customError.CustomError {
	// TODO: 校验是否有权限发布种子
	return nil
}
