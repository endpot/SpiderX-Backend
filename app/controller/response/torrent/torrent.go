package torrent

import (
	"github.com/endpot/SpiderX-Backend/app/domain/model"
	"github.com/endpot/SpiderX-Backend/app/infra/util/convert"
)

type Torrent struct {
	ID            uint        `json:"id" example:"1"`                                               // ID
	Category      interface{} `json:"category"`                                                     // 分类
	Uploader      interface{} `json:"uploader"`                                                     // 发布者
	Owner         interface{} `json:"owner"`                                                        // 拥有者
	InfoHash      string      `json:"info_hash" example:"001460cb52f9154f4ce2b510eae0dcb705210433"` // 唯一哈希码
	Title         string      `json:"title" example:"Title"`                                        // 种子标题
	SimpleDesc    string      `json:"simple_desc" example:"simple desc"`                            // 简介
	Description   string      `json:"description" example:"long long desc content"`                 // 种子介绍
	Size          uint64      `json:"size" example:"12345"`                                         // 资源体积
	SeederCount   uint        `json:"seeder_count" example:"0"`                                     // 做种者数量
	LeecherCount  uint        `json:"leecher_count" example:"0"`                                    // 下载者数量
	SnacherCount  uint        `json:"snacher_count" example:"0"`                                    // 完成都数量
	CommentCount  uint        `json:"comment_count" example:"0"`                                    // 评论数
	ViewCount     uint        `json:"view_count" example:"0"`                                       // 查看次数
	State         uint8       `json:"state" example:"0"`                                            // 种子状态 0: 待发布，1: 正常，2: 断种
	IsAnonymous   bool        `json:"is_anonymous" example:"false"`                                 // 是否匿名
	PositionLevel uint8       `json:"position_level" example:"0"`                                   // 置顶位置
	RewardBonus   uint        `json:"reward_bonus" example:"0"`                                     // 获赠魔力豆
	CreatedAt     int64       `json:"created_at" example:"1591974665"`                              // 创建时间戳（秒）
	UpdatedAt     int64       `json:"updated_at" example:"1591974665"`                              // 更新时间戳（秒）
}

type User struct {
	ID          uint   `json:"id" example:"1"`                // 用户ID
	DisplayName string `json:"display_name" example:"spider"` // 用户名
}

type Category struct {
	ID          uint   `json:"id" example:"1"`                  // 分类ID
	DisplayName string `json:"display_name" example:"document"` // 分类名
}

type Response struct {
}

func (r Response) Serialize(singleModel interface{}) interface{} {
	torrent, ok := singleModel.(*model.Torrent)
	if !ok {
		return nil
	}

	// TODO: 分类和用户相关数据补充
	return &Torrent{
		ID: torrent.ID,
		Category: &Category{
			ID:          torrent.CategoryID,
			DisplayName: "",
		},
		Uploader: &User{
			ID:          torrent.UploaderID,
			DisplayName: "",
		},
		Owner: &User{
			ID:          torrent.OwnerID,
			DisplayName: "",
		},
		InfoHash:      torrent.InfoHash,
		Title:         torrent.Title,
		SimpleDesc:    torrent.SimpleDesc,
		Description:   torrent.Description,
		Size:          torrent.Size,
		SeederCount:   torrent.SeederCount,
		LeecherCount:  torrent.LeecherCount,
		SnacherCount:  torrent.SnatcherCount,
		CommentCount:  torrent.CommentCount,
		ViewCount:     torrent.ViewCount,
		State:         torrent.State,
		IsAnonymous:   convert.ParseUint8ToBool(torrent.IsAnonymous),
		PositionLevel: torrent.PositionLevel,
		RewardBonus:   torrent.RewardBonus,
		CreatedAt:     torrent.CreatedAt.Time.Unix(),
		UpdatedAt:     torrent.UpdatedAt.Time.Unix(),
	}
}

func (r Response) Paginate(modelSlice interface{}) interface{} {
	torrentSlice, ok := modelSlice.(model.TorrentSlice)
	if !ok {
		return nil
	}

	// TODO: 逻辑补充
	resultSlice := make([]Torrent, len(torrentSlice))

	return resultSlice
}
