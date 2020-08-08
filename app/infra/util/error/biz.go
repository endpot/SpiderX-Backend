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

	// 1000~2000: Tracker
	"TRACKER__INVALID_CLIENT":  {Code: 1000, Message: "客户端不合法"},
	"TRACKER__INVALID_TORRENT": {Code: 1001, Message: "种子信息不存在"},
	"TRACKER__INVALID_USER":    {Code: 1002, Message: "用户信息不存在"},
	"TRACKER__INVALID_PARAMS":  {Code: 1003, Message: "请求参数不合法"},

	// 2000~3000: Torrent
	"TORRENT__EMPTY_TORRENT_FILE": {Code: 2000, Message: "种子文件不能为空"},
	"TORRENT__INVALID_EXTENSION":  {Code: 2001, Message: "种子文件格式错误"},
	"TORRENT__FILE_TOO_LARGE":     {Code: 2002, Message: "种子文件体积过大"},
	"TORRENT__HASH_FAILED":        {Code: 2003, Message: "种子解码失败"},
	"TORRENT__TORRENT_EXISTS":     {Code: 2004, Message: "种子已存在"},
	"TORRENT__INSERT_FAILED":      {Code: 2005, Message: "添加种子失败"},
	"TORRENT__SAVE_FAILED":        {Code: 2006, Message: "保存种子失败"},
	"TORRENT__INVALID_REQ_FORMAT": {Code: 2007, Message: "请求参数格式不正确"},
	"TORRENT__INVALID_CATEGORY":   {Code: 2008, Message: "种子分类不存在"},
	"TORRENT__INVALID_HASH":       {Code: 2009, Message: "种子已存在或未预上传种子文件"},
	"TORRENT__INVALID_POSITION":   {Code: 2010, Message: "置顶参数不正确"},
	"TORRENT__CREATE_FAILED":      {Code: 2011, Message: "创建种子失败"},

	// 2001~3000: Comment
}
