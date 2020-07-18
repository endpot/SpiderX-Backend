package error

type BizErrorStruct struct {
	Code    int
	Message string
}

var BizErrorMap = map[string]BizErrorStruct{
	// 1~1000: Default
	"DEFAULT__BAD_REQUEST":    {Code: 400, Message: "请求错误"},
	"DEFAULT__UNAUTHORIZED":   {Code: 401, Message: "需要登录"},
	"DEFAULT__FORBIDDEN":      {Code: 403, Message: "没有权限"},
	"DEFAULT__NOT_FOUND":      {Code: 404, Message: "资源不存在"},
	"DEFAULT__INTERNAL_ERROR": {Code: 500, Message: "内部错误"},

	// 1001~2000: Torrent
	"TORRENT__EMPTY_TORRENT_FILE": {Code: 1001, Message: "种子文件不能为空"},
	"TORRENT__INVALID_EXTENSION":  {Code: 1002, Message: "种子文件格式错误"},
	"TORRENT__FILE_TOO_LARGE":     {Code: 1003, Message: "种子文件体积过大"},
	"TORRENT__HASH_FAILED":        {Code: 1004, Message: "种子解码失败"},
	"TORRENT__TORRENT_EXISTS":     {Code: 1005, Message: "种子已存在"},
	"TORRENT__INSERT_FAILED":      {Code: 1006, Message: "添加种子失败"},
	"TORRENT__SAVE_FAILED":        {Code: 1007, Message: "保存种子失败"},

	// 2001~3000: Comment
}
