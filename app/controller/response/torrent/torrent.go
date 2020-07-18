package torrent

type Torrent struct {
	ID            int         `json:"id" example:"1"`                                               // ID
	Category      interface{} `json:"category"`                                                     // 分类
	Uploader      interface{} `json:"uploader"`                                                     // 发布者
	Owner         interface{} `json:"owner"`                                                        // 拥有者
	InfoHash      string      `json:"info_hash" example:"001460cb52f9154f4ce2b510eae0dcb705210433"` // 唯一哈希码
	Title         string      `json:"title" example:"Title"`                                        // 种子标题
	SimpleDesc    string      `json:"simple_desc" example:"simple desc"`                            // 简介
	Description   string      `json:"description" example:"long long desc content"`                 // 种子介绍
	Size          int         `json:"size" example:"12345"`                                         // 资源体积
	SeederCount   int         `json:"seeder_count" example:"0"`                                     // 做种者数量
	LeecherCount  int         `json:"leecher_count" example:"0"`                                    // 下载者数量
	SnacherCount  int         `json:"snacher_count" example:"0"`                                    // 完成都数量
	CommentCount  int         `json:"comment_count" example:"0"`                                    // 评论数
	ViewCount     int         `json:"view_count" example:"0"`                                       // 查看次数
	State         bool        `json:"state" example:"0"`                                            // 种子状态 0: 待发布，1: 正常，2: 断种
	IsAnonymous   bool        `json:"is_anonymous" example:"false"`                                 // 是否匿名
	PositionLevel int         `json:"position_level" example:"0"`                                   // 置顶位置
	RewardBonus   int         `json:"reward_bonus" example:"0"`                                     // 获赠魔力豆
	CreatedAt     int         `json:"created_at" example:"1591974665"`                              // 创建时间戳（秒）
	UpdatedAt     int         `json:"updated_at" example:"1591974665"`                              // 更新时间戳（秒）
}

type User struct {
	ID          int    `json:"id" example:"1"`                // 用户ID
	DisplayName string `json:"display_name" example:"spider"` // 用户名
}

type Category struct {
	ID          int    `json:"id" example:"1"`                  // 分类ID
	DisplayName string `json:"display_name" example:"document"` // 分类名
}
