package forum

// 创建论坛版块请求体
type CreateForumRequest struct {
	Name string `json:"name" validate:"required" example:"This is a forum"` // 论坛版块名
	Sort int    `json:"sort" example:"1"`                                   // 排序
}
