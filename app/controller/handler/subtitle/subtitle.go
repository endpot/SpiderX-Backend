package subtitle

import "github.com/gin-gonic/gin"

// 获取种子字幕列表
// @Summary 获取种子字幕列表
// @Description 获取种子字幕列表
// @Tags Subtitle
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param page query int false "页码" default(1)
// @Param per_page query int false "每页数量" default(15)
// @Success 200 {object} response.PageResponse{data=[]subtitle.Subtitle{uploader=subtitle.User,torrent=subtitle.Torrent}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/subtitles [get]
func GetSubtitleList(ctx *gin.Context) {
	//
}

// 获取种子字幕
// @Summary 获取种子字幕
// @Description 获取种子字幕
// @Tags Subtitle
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param subtitle_id path int true "字幕ID"
// @Success 200 {object} response.Response{data=subtitle.Subtitle{uploader=subtitle.User,torrent=subtitle.Torrent}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/subtitles/{subtitle_id} [get]
func GetSubtitle(ctx *gin.Context) {
	//
}

// 创建种子字幕
// @Summary 创建种子字幕
// @Description 创建种子字幕
// @Tags Subtitle
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param subtitle body CreateSubtitleRequest true "创建种子字幕请求参数"
// @Success 200 {object} response.Response{data=subtitle.Subtitle{uploader=subtitle.User,torrent=subtitle.Torrent}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/subtitles [post]
func CreateSubtitle(ctx *gin.Context) {
	//
}

// 预创建字幕
// @Summary 预创建字幕
// @Description 预创建字幕
// @Tags Subtitle
// @Accept multipart/form-data
// @Produce	json
// @Security ApiKeyAuth
// @Param subtitle formData file true "字幕文件"
// @Success 200 {object} response.Response{data=object{hash_code=string}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/subtitles.preUpload [post]
func PreUploadTorrent(ctx *gin.Context) {
	//
}

// 更新种子字幕
// @Summary 更新种子字幕
// @Description 更新种子字幕
// @Tags Subtitle
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param subtitle_id path int true "字幕ID"
// @Param subtitle body UpdateSubtitleRequest true "更新种子字幕请求参数"
// @Success 200 {object} response.Response{data=subtitle.Subtitle{uploader=subtitle.User,torrent=subtitle.Torrent}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/subtitles/{subtitle_id} [patch]
func UpdateSubtitle(ctx *gin.Context) {
	//
}

// 删除种子字幕
// @Summary 删除种子字幕
// @Description 删除种子字幕
// @Tags Subtitle
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param subtitle_id path int true "字幕ID"
// @Success 204 "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/subtitles/{subtitle_id} [delete]
func DeleteSubtitle(ctx *gin.Context) {
	//
}
