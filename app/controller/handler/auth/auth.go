package auth

import "github.com/gin-gonic/gin"

// 用户注册
// @Summary 用户注册
// @Description 用户自主注册
// @Tags Auth
// @Accept json
// @Produce	json
// @Param auth body RegisterRequest true "用户注册参数"
// @Success 204 "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /auth.register [post]
func Register(c *gin.Context) {
	//
}

// 用户登录
// @Summary 用户登录
// @Description 用户使用用户名密码登录
// @Tags Auth
// @Accept json
// @Produce	json
// @Param auth body LoginRequest true "用户登录参数"
// @Success 200 {object} response.Response{data=object{token=string}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /auth.login [post]
func Login(c *gin.Context) {
	//
}

// 用户登出
// @Summary 用户登出
// @Description 用户退出登录态
// @Tags Auth
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Success 204 "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /auth.logout [post]
func Logout(c *gin.Context) {
	//
}

// 重置密码
// @Summary 重置密码
// @Description 重置密码
// @Tags Auth
// @Accept json
// @Produce	json
// @Param auth body ResetRequest true "重置密码参数"
// @Success 204 "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /auth.reset [post]
func Reset(c *gin.Context) {
	//
}

// 获取个人信息
// @Summary 获取个人信息
// @Description 获取个人信息
// @Tags Auth
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=auth.Mine} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /auth.getMine [post]
func GetMine(c *gin.Context) {
	//
}

// 发送校验码
// @Summary 发送校验码
// @Description 发送校验码
// @Tags Auth
// @Accept json
// @Produce	json
// @Param auth body SendCodeRequest true "发送校验码参数"
// @Success 204 "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /auth.sendCode [post]
func SendCode(c *gin.Context) {
	//
}

// 刷新登录密钥
// @Summary 刷新登录密钥
// @Description 刷新登录密钥
// @Tags Auth
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=object{token=string}} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /auth.refreshToken [post]
func RefreshToken(c *gin.Context) {
	//
}
