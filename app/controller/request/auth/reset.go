package auth

// 重置密码请求体
type ResetRequest struct {
	Email           string `json:"email" validate:"required" example:"spider@spider.com"`                         // 邮箱
	Password        string `json:"password" validate:"required" minLength:"8" maxLength:"32" example:"123456789"` // 密码
	PasswordConfirm string `json:"password_confirm" validate:"required" example:"123456789"`                      // 密码确认
	VerifyCode      string `json:"verify_code" example:"spider"`                                                  // 校验码
}
