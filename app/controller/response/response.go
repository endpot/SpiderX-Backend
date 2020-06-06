package response

type Response struct {
	Code int         `json:"code" example:"200"` // 业务状态码
	Data interface{} `json:"data"`               // 数据
}

type ErrResponse struct {
	Code    int    `json:"code" example:"500"` // 业务状态码
	Message string `json:"message"`            // 消息
}

type PageMeta struct {
	Page    int `json:"page" example:"1"`      // 页码
	PerPage int `json:"per_page" example:"15"` // 每页数量
	Total   int `json:"total" example:"100"`   // 总数
}

type PageResponse struct {
	Code int         `json:"code" example:"200"` // 业务状态码
	Data interface{} `json:"data"`               // 数据
	Meta PageMeta    `json:"meta"`               // 附加信息
}

func NewResponse(code int, data interface{}) *Response {
	return &Response{
		Code: code,
		Data: data,
	}
}

func NewErrResponse(code int, message string) *ErrResponse {
	return &ErrResponse{
		Code:    code,
		Message: message,
	}
}

func NewPageResponse(code int, data interface{}, page int, perPage int, total int) *PageResponse {
	return &PageResponse{
		Code: code,
		Data: data,
		Meta: PageMeta{
			Page:    page,
			PerPage: perPage,
			Total:   total,
		},
	}
}
