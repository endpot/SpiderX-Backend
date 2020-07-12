package post

import "github.com/gin-gonic/gin"

// 获取主题回复列表
// @Summary 获取主题回复列表
// @Description 获取主题回复列表
// @Tags Post
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param topic_id path int true "主题ID"
// @Param page query int false "页码" default(1)
// @Param per_page query int false "每页数量" default(15)
// @Success 200 {object} response.PageResponse{data=[]post.Post{user=post.User,topic=post.Topic}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /topics/{topic_id}/posts [get]
func GetPostList(ctx *gin.Context) {
	//
}

// 获取主题回复
// @Summary 获取主题回复
// @Description 获取主题回复
// @Tags Post
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param topic_id path int true "主题ID"
// @Param post_id path int true "回复ID"
// @Success 200 {object} response.Response{data=post.Post{user=post.User,topic=post.Topic}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /topics/{topic_id}/posts/{post_id} [get]
func GetPost(ctx *gin.Context) {
	//
}

// 创建回复
// @Summary 创建回复
// @Description 创建回复
// @Tags Post
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param topic_id path int true "主题ID"
// @Param post body CreatePostRequest true "创建回复请求参数"
// @Success 200 {object} response.Response{data=post.Post{user=post.User,topic=post.Topic}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /topics/{topic_id}/posts [post]
func CreatePost(ctx *gin.Context) {
	//
}

// 更新回复
// @Summary 更新回复
// @Description 更新回复
// @Tags Post
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param topic_id path int true "主题ID"
// @Param post_id path int true "回复ID"
// @Param post body UpdatePostRequest true "更新回复请求参数"
// @Success 200 {object} response.Response{data=post.Post{user=post.User,topic=post.Topic}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /topics/{topic_id}/posts/{post_id} [patch]
func UpdatePost(ctx *gin.Context) {
	//
}

// 删除回复
// @Summary 删除回复
// @Description 删除回复
// @Tags Post
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param topic_id path int true "主题ID"
// @Param post_id path int true "回复ID"
// @Success 204 "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /topics/{topic_id}/posts/{post_id} [delete]
func DeletePost(ctx *gin.Context) {
	//
}
