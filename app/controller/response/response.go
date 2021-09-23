package response

type IResponse interface {
	Serialize(singleModel interface{}) interface{}
	Paginate(modelSlice interface{}) interface{}
}

type Response struct {
	Code int         `json:"code" example:"0"` // 业务状态码
	Data interface{} `json:"data"`             // 数据
}

type ErrResponse struct {
	Code    int    `json:"code" example:"400"` // 业务错误码
	Message string `json:"message"`            // 业务错误信息
}

type PageMeta struct {
	Page    int `json:"page" example:"1"`      // 页码
	PerPage int `json:"per_page" example:"15"` // 每页数量
	Total   int `json:"total" example:"100"`   // 总数
}

type PageResponse struct {
	Code int         `json:"code" example:"0"` // 业务状态码
	Data interface{} `json:"data"`             // 数据
	Meta PageMeta    `json:"meta"`             // 附加信息
}

type Serializer struct {
}

var RespSerializer *Serializer

func init() {
	RespSerializer = &Serializer{}
}

func (serializer *Serializer) Serialize(resp IResponse, singleModel interface{}) *Response {
	return &Response{
		Code: 0,
		Data: resp.Serialize(singleModel),
	}
}

func (serializer *Serializer) Paginate(resp IResponse, modelSlice interface{}, page int, perPage int) *PageResponse {
	return &PageResponse{
		Code: 0,
		Data: resp.Paginate(modelSlice),
		Meta: PageMeta{
			Page:    page,
			PerPage: perPage,
		},
	}
}
