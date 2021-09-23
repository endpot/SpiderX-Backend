package post

// 更新主题回复请求体
type UpdatePostRequest struct {
	Content string `json:"content" example:"This is a content"` // 回复内容
}
