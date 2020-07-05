package torrent

// 创建种子请求体
type CreateTorrentRequest struct {
	CategoryId    int    `json:"category_id" validate:"required" example:"401"`                                    // 分类ID
	InfoHash      string `json:"info_hash" validate:"required" example:"001460cb52f9154f4ce2b510eae0dcb705210433"` // 唯一哈希码
	Title         string `json:"title" validate:"required" example:"Title"`                                        // 种子标题
	SimpleDesc    string `json:"simple_desc" example:"simple desc"`                                                // 简介
	Description   string `json:"description" validate:"required" example:"long long desc content"`                 // 种子介绍
	IsAnonymous   bool   `json:"is_anonymous" default:"false" example:"false"`                                     // 是否匿名
	PositionLevel int    `json:"position_level" default:"0" example:"0"`                                           // 置顶位置
}
