package topic

import "github.com/gin-gonic/gin"

// 获取主题列表
// @Summary 获取主题列表
// @Description 获取主题列表
// @Tags Topic
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param forum_id path int true "论坛版块ID"
// @Param page query int false "页码" default(1)
// @Param per_page query int false "每页数量" default(15)
// @Success 200 {object} response.PageResponse{data=[]topic.Topic{user=topic.User,forum=topic.Forum}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums/{forum_id}/topics [get]
func GetTopicList(c *gin.Context) {
	//
}

// 获取主题信息
// @Summary 获取主题信息
// @Description 获取主题信息
// @Tags Topic
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param forum_id path int true "论坛版块ID"
// @Success 200 {object} response.Response{data=topic.Topic{user=topic.User,forum=topic.Forum}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums/{forum_id}/topics/{topic_id} [get]
func GetTopic(c *gin.Context) {
	//
}

// 创建主题
// @Summary 创建主题
// @Description 创建主题
// @Tags Topic
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param forum_id path int true "论坛版块ID"
// @Param topic body CreateTopicRequest true "创建主题请求参数"
// @Success 200 {object} response.Response{data=topic.Topic{user=topic.User,forum=topic.Forum}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums/{forum_id}/topics [post]
func CreateTopic(c *gin.Context) {
	//
}

// 更新主题
// @Summary 更新主题
// @Description 更新主题
// @Tags Topic
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param forum_id path int true "论坛版块ID"
// @Param topic_id path int true "主题ID"
// @Param topic body UpdateTopicRequest true "更新主题请求参数"
// @Success 200 {object} response.Response{data=topic.Topic{user=topic.User,forum=topic.Forum}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums/{forum_id}/topics/{topic_id} [patch]
func UpdateTopic(c *gin.Context) {
	//
}

// 删除主题
// @Summary 删除主题
// @Description 删除主题
// @Tags Topic
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param forum_id path int true "论坛版块ID"
// @Param topic_id path int true "主题ID"
// @Success 204 "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /forums/{forum_id}/topics/{topic_id} [delete]
func DeleteTopic(c *gin.Context) {
	//
}
