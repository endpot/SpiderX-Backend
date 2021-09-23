package notice

import "github.com/gin-gonic/gin"

// 获取站点公告列表
// @Summary 获取站点公告列表
// @Description 获取站点公告列表
// @Tags Notice
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param page query int false "页码" default(1)
// @Param per_page query int false "每页数量" default(15)
// @Success 200 {object} response.PageResponse{data=[]notice.Notice} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /notices [get]
func GetNoticeList(ctx *gin.Context) {
	//
}

// 获取站点公告
// @Summary 获取站点公告
// @Description 获取站点公告
// @Tags Notice
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param notice_id path int true "公告ID"
// @Success 200 {object} response.Response{data=notice.Notice} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /notices/{notice_id} [get]
func GetNotice(ctx *gin.Context) {
	//
}

// 创建公告
// @Summary 创建公告
// @Description 创建公告
// @Tags Notice
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param notice body CreateNoticeRequest true "创建公告请求参数"
// @Success 200 {object} response.Response{data=notice.Notice} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /notices [post]
func CreateNotice(ctx *gin.Context) {
	//
}

// 更新公告
// @Summary 更新公告
// @Description 更新公告
// @Tags Notice
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param notice_id path int true "公告ID"
// @Param notice body UpdateNoticeRequest true "更新公告请求参数"
// @Success 200 {object} response.Response{data=notice.Notice} "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /notices/{notice_id} [patch]
func UpdateNotice(ctx *gin.Context) {
	//
}

// 删除公告
// @Summary 删除公告
// @Description 删除公告
// @Tags Notice
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param notice_id path int true "公告ID"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.ErrResponse "请求参数异常"
// @Failure 401 {object} response.ErrResponse "用户身份信息异常"
// @Failure 403 {object} response.ErrResponse "没有操作权限"
// @Failure 404 {object} response.ErrResponse "没有对象"
// @Failure 500 {object} response.ErrResponse "内部错误"
// @Router /notices/{notice_id} [delete]
func DeleteNotice(ctx *gin.Context) {
	//
}
