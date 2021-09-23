package subtitle

// 更新种子评论请求体
type UpdateSubtitleRequest struct {
	Title       string `json:"title" example:"Title"`                     // 标题
	Language    string `json:"language" validate:"required" example:"en"` // 语言
	IsAnonymous bool   `json:"is_anonymous" example:"false"`              // 是否匿名
}
