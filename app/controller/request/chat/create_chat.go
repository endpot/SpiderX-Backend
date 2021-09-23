package chat

// 创建群聊消息请求体
type CreateChatRequest struct {
	Content string `json:"content" validate:"required" example:"This is a chat"` // 聊天内容
}
