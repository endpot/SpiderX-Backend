package forum

type Forum struct {
	ID        int    `json:"id" example:"1"`                      // ID
	Name      string `json:"name" example:"This is a forum name"` // 版块名称
	Sort      int    `json:"sort" example:"1"`                    // 排序
	CreatedAt int    `json:"created_at" example:"1591974665"`     // 创建时间戳（秒）
	UpdatedAt int    `json:"updated_at" example:"1591974665"`     // 更新时间戳（秒）
}
