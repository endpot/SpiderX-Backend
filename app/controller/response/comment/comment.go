package comment

type Comment struct {
	ID        int         `json:"id" example:"1"`                      // ID
	User      interface{} `json:"user"`                                // 用户
	Torrent   interface{} `json:"torrent"`                             // 种子
	Content   string      `json:"content" example:"This is a content"` // 回复内容
	CreatedAt int         `json:"created_at" example:"1591974665"`     // 创建时间戳（秒）
	UpdatedAt int         `json:"updated_at" example:"1591974665"`     // 更新时间戳（秒）
}

type User struct {
	ID          int    `json:"id" example:"1"`                // 用户ID
	DisplayName string `json:"display_name" example:"spider"` // 用户名
}

type Torrent struct {
	ID    int    `json:"id" example:"1"`                  // 种子ID
	Title string `json:"title" example:"This is a title"` // 种子标题
}
