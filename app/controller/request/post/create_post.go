package post

// 创建主题回复请求体
type CreatePostRequest struct {
	Content string `json:"content" validate:"required" example:"This is a content"` // 回复内容
}
