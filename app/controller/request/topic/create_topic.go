package topic

// 创建主题请求体
type CreateTopicRequest struct {
	Title    string `json:"title" validate:"required" example:"This is a title"`     // 主题标题
	Content  string `json:"content" validate:"required" example:"This is a content"` // 主题内容
	Position int    `json:"position" default:"0" example:"0"`                        // 置顶位置
}
