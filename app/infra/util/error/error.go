package error

import (
	"encoding/json"
	"github.com/endpot/SpiderX-Backend/app/infra/util/http"
	"github.com/gin-gonic/gin"
)

type CustomError struct {
	Status  int    // HTTP 状态
	Code    int    // 业务错误码
	Message string // 业务错误信息
}

func (e *CustomError) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func (e *CustomError) Serialize() map[string]interface{} {
	return gin.H{
		"code":    e.Code,
		"message": e.Message,
	}
}

func NewCustomError(status int, code int, message string) *CustomError {
	return &CustomError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func NewBadRequestError(errCode string) *CustomError {
	if val, ok := BizErrorMap[errCode]; ok {
		return NewCustomError(http.StatusBadRequest, val.Code, val.Message)
	}

	return NewBadRequestError("DEFAULT__BAD_REQUEST")
}

func NewUnauthorizedError(errCode string) *CustomError {
	if val, ok := BizErrorMap[errCode]; ok {
		return NewCustomError(http.StatusUnauthorized, val.Code, val.Message)
	}

	return NewUnauthorizedError("DEFAULT__UNAUTHORIZED")
}

func NewForbiddenError(errCode string) *CustomError {
	if val, ok := BizErrorMap[errCode]; ok {
		return NewCustomError(http.StatusForbidden, val.Code, val.Message)
	}

	return NewForbiddenError("DEFAULT__FORBIDDEN")
}

func NewNotFoundError(errCode string) *CustomError {
	if val, ok := BizErrorMap[errCode]; ok {
		return NewCustomError(http.StatusNotFound, val.Code, val.Message)
	}

	return NewNotFoundError("DEFAULT__NOT_FOUND")
}

func NewInternalServerError(errCode string) *CustomError {
	if val, ok := BizErrorMap[errCode]; ok {
		return NewCustomError(http.StatusInternalServerError, val.Code, val.Message)
	}

	return NewInternalServerError("DEFAULT__INTERNAL_ERROR")
}
