package comment

// 更新种子评论请求体
type UpdateCommentRequest struct {
	Content string `json:"content" example:"This is a content"` // 评论内容
}
