package auth

type Mine struct {
	ID          int      `json:"id" example:"1"`                                  // ID
	DisplayName string   `json:"display_name" example:"spider"`                   // 用户名
	Email       string   `json:"email" example:"spider@spider.com"`               // 用户邮箱
	Avatar      string   `json:"avatar" example:"https://i.endpot.com/x/xxx.jpg"` // 头像地址
	Grade       int      `json:"grade" example:"1"`                               // 用户等级
	Roles       []string `json:"roles"`                                           // 用户角色
	UpTraffic   int      `json:"up_traffic" example:"1"`                          // 上传量
	DownTraffic int      `json:"down_traffic" example:"1"`                        // 下载量
}
