package topic

type Topic struct {
	ID        int         `json:"id" example:"1"`                      // ID
	User      interface{} `json:"user"`                                // 用户
	Forum     interface{} `json:"forum"`                               // 论坛版块
	Title     string      `json:"title" example:"This is a title"`     // 主题标题
	Content   string      `json:"content" example:"This is a content"` // 主题内容
	ViewCount int         `json:"view_count" example:"0"`              // 查看次数
	Position  int         `json:"position" example:"0"`                // 置顶位置
	IsHot     bool        `json:"is_hot" example:"false"`              // 是否热门
	CreatedAt int         `json:"created_at" example:"1591974665"`     // 创建时间戳（秒）
	UpdatedAt int         `json:"updated_at" example:"1591974665"`     // 更新时间戳（秒）
}

type User struct {
	ID          int    `json:"id" example:"1"`                // 用户ID
	DisplayName string `json:"display_name" example:"spider"` // 用户名
}

type Forum struct {
	ID   int    `json:"id" example:"1"`                      // 版块ID
	Name string `json:"name" example:"This is a forum name"` // 版块名称
}
