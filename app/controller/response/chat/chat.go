package chat

type Chat struct {
	ID        int         `json:"id" example:"1"`                     // ID
	User      interface{} `json:"user"`                               // 用户
	Content   string      `json:"content" example:"This is a notice"` // 聊天内容
	CreatedAt int         `json:"created_at" example:"1591974665"`    // 创建时间戳（秒）
	UpdatedAt int         `json:"updated_at" example:"1591974665"`    // 更新时间戳（秒）
}

type User struct {
	ID          int    `json:"id" example:"1"`                // 用户ID
	DisplayName string `json:"display_name" example:"spider"` // 用户名
}
