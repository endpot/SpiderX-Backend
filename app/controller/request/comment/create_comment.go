package comment

// 创建种子评论请求体
type CreateCommentRequest struct {
	Content string `json:"content" validate:"required" example:"This is a content"` // 评论内容
}
