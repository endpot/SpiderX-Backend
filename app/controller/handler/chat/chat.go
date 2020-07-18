package chat

import "github.com/gin-gonic/gin"

// 获取群聊消息列表
// @Summary 获取群聊消息列表
// @Description 获取群聊消息列表
// @Tags Chat
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param page query int false "页码" default(1)
// @Param per_page query int false "每页数量" default(15)
// @Success 200 {object} response.PageResponse{data=[]chat.Chat{user=chat.User}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /chats [get]
func GetChatList(ctx *gin.Context) {
	//
}

// 获取群聊消息
// @Summary 获取群聊消息
// @Description 获取群聊消息
// @Tags Chat
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param chat_id path int true "群聊ID"
// @Success 200 {object} response.Response{data=chat.Chat{user=chat.User}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /chats/{chat_id} [get]
func GetChat(ctx *gin.Context) {
	//
}

// 创建一条群聊消息
// @Summary 创建群聊
// @Description 创建群聊
// @Tags Chat
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param chat body CreateChatRequest true "创建聊天消息请求参数"
// @Success 200 {object} response.Response{data=chat.Chat{user=chat.User}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /chats [post]
func CreateChat(ctx *gin.Context) {
	//
}

// 更新群聊消息
// @Summary 更新群聊消息
// @Description 更新群聊消息
// @Tags Chat
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param chat_id path int true "群聊ID"
// @Param chat body UpdateChatRequest true "更新群聊请求参数"
// @Success 200 {object} response.Response{data=chat.Chat{user=chat.User}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /chats/{chat_id} [patch]
func UpdateChat(ctx *gin.Context) {
	//
}

// 删除群聊消息
// @Summary 删除群聊消息
// @Description 删除群聊消息
// @Tags Chat
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param chat_id path int true "群聊ID"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /chats/{chat_id} [delete]
func DeleteChat(ctx *gin.Context) {
	//
}
