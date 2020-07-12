package forum

import "github.com/gin-gonic/gin"

// 获取论坛版块列表
// @Summary 获取论坛版块列表
// @Description 获取论坛版块列表
// @Tags Forum
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param page query int false "页码" default(1)
// @Param per_page query int false "每页数量" default(15)
// @Success 200 {object} response.PageResponse{data=[]forum.Forum{}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums [get]
func GetForumList(ctx *gin.Context) {
	//
}

// 获取论坛版块信息
// @Summary 获取论坛版块信息
// @Description 获取论坛版块信息
// @Tags Forum
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param forum_id path int true "论坛版块ID"
// @Success 200 {object} response.Response{data=forum.Forum{}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums/{forum_id} [get]
func GetForum(ctx *gin.Context) {
	//
}

// 创建论坛版块
// @Summary 创建论坛版块
// @Description 创建论坛版块
// @Tags Forum
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param forum body CreateForumRequest true "创建论坛版块请求参数"
// @Success 200 {object} response.Response{data=forum.Forum{}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums [post]
func CreateForum(ctx *gin.Context) {
	//
}

// 更新论坛版块
// @Summary 更新论坛版块
// @Description 更新论坛版块
// @Tags Forum
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param forum_id path int true "论坛版块ID"
// @Param forum body UpdateForumRequest true "更新论坛版块请求参数"
// @Success 200 {object} response.Response{data=forum.Forum{}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums/{forum_id} [patch]
func UpdateForum(ctx *gin.Context) {
	//
}

// 删除论坛版块
// @Summary 删除论坛版块
// @Description 删除论坛版块
// @Tags Forum
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param forum_id path int true "论坛版块ID"
// @Success 204 "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums/{forum_id} [delete]
func DeleteForum(ctx *gin.Context) {
	//
}
