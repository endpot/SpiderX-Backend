package auth

// 用户登录请求体
type LoginRequest struct {
	DisplayName string `json:"display_name" validate:"required" example:"spider"` // 用户名
	Password    string `json:"password" validate:"required" example:"123456789"`  // 密码
}
