package notice

// 更新公告请求体
type UpdateNoticeRequest struct {
	Title   string `json:"title" validate:"required" example:"title"`              // 公告标题
	Content string `json:"content" validate:"required" example:"This is a notice"` // 公告内容
}
