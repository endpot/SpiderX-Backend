package torrent

import (
	"github.com/endpot/SpiderX-Backend/app/domain/model"
	torrentExt "github.com/endpot/SpiderX-Backend/app/domain/modext/torrent"
	"github.com/endpot/SpiderX-Backend/app/infra/db"
	customError "github.com/endpot/SpiderX-Backend/app/infra/util/error"
	"github.com/gin-gonic/gin"
)

// 创建种子请求，可以分为 uri、form、body 三部分，其中 body 统一使用 json 格式
// 使用 gin 的 binding 标明 required 属性
// 对于必填项，要定义 example
// 对于非必填项，要指明 default
type CreateTorrentRequest struct {
	CreateTorrentRequestBody
}

// 请求的 Body 部分
type CreateTorrentRequestBody struct {
	CategoryID    uint   `json:"category_id" example:"401" binding:"required,number"`              // 分类ID
	InfoHash      string `json:"info_hash" binding:"required,len=40,alphanum"`                     // 唯一哈希码
	Title         string `json:"title" maxLength:"255" example:"Title" binding:"required,lte=255"` // 种子标题
	SimpleDesc    string `json:"simple_desc" example:"simple desc" default:""`                     // 简介
	Description   string `json:"description" example:"long long desc content" bind:"required"`     // 种子介绍
	IsAnonymous   bool   `json:"is_anonymous" default:"false"`                                     // 是否匿名
	PositionLevel uint8  `json:"position_level" default:"0" binding:"number"`                      // 置顶位置
}

func (request *CreateTorrentRequest) ValidateRequest(ctx *gin.Context) *customError.CustomError {
	// 校验参数格式
	if err := ctx.ShouldBind(request); err != nil {
		return customError.NewBadRequestError("TORRENT__INVALID_REQ_FORMAT")
	}

	// 校验参数业务合法性
	{
		// TODO: 通过对比数据库校验 CategoryID 是否合法
		if request.CategoryID <= 0 {
			return customError.NewBadRequestError("TORRENT__INVALID_CATEGORY")
		}

		// 校验 InfoHash 是否合法
		if _, err := model.Torrents(
			model.TorrentWhere.InfoHash.EQ(request.InfoHash),
			model.TorrentWhere.State.EQ(torrentExt.PENDING_STATE),
		).One(ctx, db.DB); err != nil {
			return customError.NewBadRequestError("TORRENT__INVALID_HASH")
		}

		// 校验 PositionLevel 是否合法
		if request.PositionLevel < 0 || request.PositionLevel > torrentExt.LAST_STICKY_POSITION {
			return customError.NewBadRequestError("TORRENT__INVALID_POSITION")
		}
	}

	return nil
}

func (request *CreateTorrentRequest) ValidateAuth(ctx *gin.Context) *customError.CustomError {
	// TODO: 校验是否有权限发布种子

	// TODO: 校验是否有权限置顶
	return nil
}
