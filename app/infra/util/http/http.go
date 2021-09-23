package http

import "net/http"

// 本系统仅使用常用 HTTP 状态码
const (
	StatusOK = http.StatusOK // 正常响应

	StatusBadRequest   = http.StatusBadRequest   // 请求不合法
	StatusUnauthorized = http.StatusUnauthorized // 缺少有效的登录信息
	StatusForbidden    = http.StatusForbidden    // 没有操作权限
	StatusNotFound     = http.StatusNotFound     // 操作资源不存在

	StatusInternalServerError = http.StatusInternalServerError // 内部错误
)
