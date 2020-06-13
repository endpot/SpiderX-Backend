package notice

type Notice struct {
	ID        int    `json:"id" example:"1"`                     // ID
	Title     string `json:"title" example:"title"`              // 公告标题
	Content   string `json:"content" example:"This is a notice"` // 公告内容
	CreatedAt int    `json:"created_at" example:"1591974665"`    // 创建时间戳（秒）
	UpdatedAt int    `json:"updated_at" example:"1591974665"`    // 更新时间戳（秒）
}
