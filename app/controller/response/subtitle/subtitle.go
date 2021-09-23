package subtitle

type Subtitle struct {
	ID            int         `json:"id" example:"1"`                  // ID
	HashCode      string      `json:"hash_code"`                       // HASH码
	Torrent       interface{} `json:"torrent"`                         // 种子
	Uploader      interface{} `json:"uploader"`                        // 上传者
	Title         string      `json:"title" example:"Title"`           // 标题
	Filename      string      `json:"filename" example:"abc.zip"`      // 文件名
	Language      string      `json:"language" example:"en"`           // 语言
	Extension     string      `json:"extension" example:"zip"`         // 后缀
	DownloadCount int         `json:"download_count" example:"0"`      // 下载次数
	IsAvailable   bool        `json:"is_available" example:"false"`    // 就否可用
	IsAnonymous   bool        `json:"is_anonymous" example:"false"`    // 是否匿名
	CreatedAt     int         `json:"created_at" example:"1591974665"` // 创建时间戳（秒）
	UpdatedAt     int         `json:"updated_at" example:"1591974665"` // 更新时间戳（秒）
}

type User struct {
	ID          int    `json:"id" example:"1"`                // 用户ID
	DisplayName string `json:"display_name" example:"spider"` // 用户名
}

type Torrent struct {
	ID    int    `json:"id" example:"1"`        // 种子ID
	Title string `json:"title" example:"title"` // 种子标题
}
