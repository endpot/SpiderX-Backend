package torrent

// 更新种子请求体
type UpdateTorrentRequest struct {
	CategoryId    int    `json:"category_id"`                                  // 分类ID
	Title         string `json:"title" example:"Title"`                        // 种子标题
	SimpleDesc    string `json:"simple_desc" example:"simple desc"`            // 简介
	Description   string `json:"description" example:"long long desc content"` // 种子介绍
	State         int    `json:"state" example:"0"`                            // 种子状态 0: 待发布，1: 正常，2: 断种
	IsAnonymous   bool   `json:"is_anonymous" example:"false"`                 // 是否匿名
	PositionLevel int    `json:"position_level" example:"0"`                   // 置顶位置
}
