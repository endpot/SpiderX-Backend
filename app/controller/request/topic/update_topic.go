package topic

// 更新主题请求体
type UpdateTopicRequest struct {
	Title    string `json:"title" example:"This is a title"`     // 主题标题
	Content  string `json:"content" example:"This is a content"` // 主题内容
	Position int    `json:"position" example:"0"`                // 置顶位置
}
