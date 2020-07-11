package subtitle

// 创建字幕请求体
type CreateSubtitleRequest struct {
	HashCode    string `json:"hash_code" validate:"required"`                // HASH码
	Title       string `json:"title" example:"Title"`                        // 标题
	Language    string `json:"language" validate:"required" example:"en"`    // 语言
	IsAnonymous bool   `json:"is_anonymous" default:"false" example:"false"` // 是否匿名
}
