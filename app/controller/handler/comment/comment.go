package comment

import "github.com/gin-gonic/gin"

// 获取种子评论列表
// @Summary 获取种子评论列表
// @Description 获取种子评论列表
// @Tags Comment
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param page query int false "页码" default(1)
// @Param per_page query int false "每页数量" default(15)
// @Success 200 {object} response.PageResponse{data=[]comment.Comment{user=comment.User,torrent=comment.Torrent}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/comments [get]
func GetCommentList(ctx *gin.Context) {
	//
}

// 获取种子评论
// @Summary 获取种子评论
// @Description 获取种子评论
// @Tags Comment
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param comment_id path int true "评论ID"
// @Success 200 {object} response.Response{data=comment.Comment{user=comment.User,torrent=comment.Torrent}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/comments/{comment_id} [get]
func GetComment(ctx *gin.Context) {
	//
}

// 创建种子评论
// @Summary 创建种子评论
// @Description 创建种子评论
// @Tags Comment
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param comment body CreateCommentRequest true "创建种子评论请求参数"
// @Success 200 {object} response.Response{data=comment.Comment{user=comment.User,torrent=comment.Torrent}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/comments [post]
func CreateComment(ctx *gin.Context) {
	//
}

// 更新种子评论
// @Summary 更新种子评论
// @Description 更新种子评论
// @Tags Comment
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param comment_id path int true "评论ID"
// @Param comment body UpdateCommentRequest true "更新种子评论请求参数"
// @Success 200 {object} response.Response{data=comment.Comment{user=comment.User,torrent=comment.Torrent}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/comments/{comment_id} [patch]
func UpdateComment(ctx *gin.Context) {
	//
}

// 删除种子评论
// @Summary 删除种子评论
// @Description 删除种子评论
// @Tags Comment
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param torrent_id path int true "种子ID"
// @Param comment_id path int true "评论ID"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /torrents/{torrent_id}/comments/{comment_id} [delete]
func DeleteComment(ctx *gin.Context) {
	//
}
