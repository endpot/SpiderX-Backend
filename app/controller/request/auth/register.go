package auth

// 用户注册请求体
type RegisterRequest struct {
	DisplayName     string `json:"display_name" validate:"required" minLength:"1" maxLength:"32" example:"spider"` // 用户名
	Email           string `json:"email" validate:"required" example:"spider@spider.com"`                          // 邮箱
	Password        string `json:"password" validate:"required" minLength:"8" maxLength:"32" example:"123456789"`  // 密码
	PasswordConfirm string `json:"password_confirm" validate:"required" example:"123456789"`                       // 密码确认
	InviteCode      string `json:"invite_code" example:"spider"`                                                   // 邀请码
	VerifyCode      string `json:"verify_code" example:"spider"`                                                   // 校验码
}
