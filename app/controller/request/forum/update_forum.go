package forum

// 更新论坛版块请求体
type UpdateForumRequest struct {
	Name string `json:"name" validate:"required" example:"This is a forum"` // 论坛版块名
	Sort int    `json:"sort" example:"1"`                                   // 排序
}
