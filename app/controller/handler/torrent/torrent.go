package torrent

import (
	"github.com/endpot/SpiderX-Backend/app/controller/request"
	torrentRequest "github.com/endpot/SpiderX-Backend/app/controller/request/torrent"
	"github.com/endpot/SpiderX-Backend/app/controller/response"
	torrentResonse "github.com/endpot/SpiderX-Backend/app/controller/response/torrent"
	torrentService "github.com/endpot/SpiderX-Backend/app/domain/service/torrent"
	"github.com/endpot/SpiderX-Backend/app/infra/util/http"
	"github.com/gin-gonic/gin"
)

// 获取种子列表
// @Summary 获取种子列表
// @Description 获取种子列表
// @Tags Torrent
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param page query int false "页码" default(1)
// @Param per_page query int false "每页数量" default(15)
// @Success 200 {object} response.PageResponse{data=[]torrent.Torrent{category=torrent.Category,uploader=torrent.User,owner=torrent.User}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents [get]
func GetTorrentList(ctx *gin.Context) {
	//
}

// 获取种子信息
// @Summary 获取种子信息
// @Description 获取种子信息
// @Tags Torrent
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Success 200 {object} response.Response{data=torrent.Torrent{category=torrent.Category,uploader=torrent.User,owner=torrent.User}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id} [get]
func GetTorrent(ctx *gin.Context) {
	//
}

// 创建种子
// @Summary 创建种子
// @Description 创建种子
// @Tags Torrent
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent body CreateTorrentRequestBody true "创建种子请求参数"
// @Success 200 {object} response.Response{data=torrent.Torrent{category=torrent.Category,uploader=torrent.User,owner=torrent.User}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents [post]
func CreateTorrent(ctx *gin.Context) {
	// 使用 CreateTorrentRequest 校验请求
	req := &torrentRequest.CreateTorrentRequest{}
	if _, err := request.Validator.Validate(
		ctx,
		req,
	); err != nil {
		ctx.JSON(err.Status, err.Serialize())
		return
	}

	// 调用 CreateTorrent 服务
	torrent, err := torrentService.CreateTorrent(ctx, req)
	if err != nil {
		ctx.JSON(err.Status, err.Serialize())
		return
	}

	resp := &torrentResonse.Response{}
	ctx.JSON(http.StatusOK, response.RespSerializer.Serialize(
		resp, torrent,
	))
}

// 预创建种子
// @Summary 预创建种子
// @Description 预创建种子
// @Tags Torrent
// @Accept multipart/form-data
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent formData file true "种子文件"
// @Success 200 {object} response.Response{data=object{info_hash=string}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents.preUpload [post]
func PreUploadTorrent(ctx *gin.Context) {
	// 使用 PreUploadTorrentRequest 校验请求
	req := &torrentRequest.PreUploadTorrentRequest{}
	if _, err := request.Validator.Validate(
		ctx,
		req,
	); err != nil {
		ctx.JSON(err.Status, err.Serialize())
		return
	}

	// 调用 PreUploadTorrent 服务
	torrent, err := torrentService.PreUploadTorrent(ctx, req)
	if err != nil {
		ctx.JSON(err.Status, err.Serialize())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"info_hash": torrent.InfoHash,
	})
}

// 更新种子
// @Summary 更新种子
// @Description 更新种子
// @Tags Torrent
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param torrent body UpdateTorrentRequest true "更新种子请求参数"
// @Success 200 {object} response.Response{data=torrent.Torrent{category=torrent.Category,uploader=torrent.User,owner=torrent.User}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id} [patch]
func UpdateTorrent(ctx *gin.Context) {
	//
}

// 删除种子
// @Summary 删除种子
// @Description 删除种子
// @Tags Torrent
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id} [delete]
func DeleteTorrent(ctx *gin.Context) {
	//
}
