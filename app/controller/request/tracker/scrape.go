package tracker

import (
	customError "github.com/endpot/SpiderX-Backend/app/infra/util/error"
	"github.com/gin-gonic/gin"
)

type ScrapeRequest struct {
	InfoHashSlice []string `form:"info_hash" binding:"required"` // 唯一哈希码
}

func (request *ScrapeRequest) ValidateRequest(ctx *gin.Context) *customError.CustomError {
	// TODO: 客户端白名单校验

	// 校验参数格式
	if err := ctx.ShouldBind(request); err != nil {
		return customError.NewBadRequestError("TORRENT__INVALID_REQ_FORMAT")
	}

	// 校验参数业务合法性
	{
		//
	}

	return nil
}

func (request *ScrapeRequest) ValidateAuth(ctx *gin.Context) *customError.CustomError {
	return nil
}
