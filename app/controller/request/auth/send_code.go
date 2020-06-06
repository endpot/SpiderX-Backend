package auth

// 发送校验码请求体
type SendCodeRequest struct {
	Email string `json:"email" validate:"required" example:"spider@spider.com"` // 用户邮箱
}
