package request

import (
	customError "github.com/endpot/SpiderX-Backend/app/infra/util/error"
	"github.com/gin-gonic/gin"
)

type IRequest interface {
	ValidateRequest(ctx *gin.Context) *customError.CustomError
	ValidateAuth(ctx *gin.Context) *customError.CustomError
}

type ReqValidator struct {
}

func (validator *ReqValidator) Validate(ctx *gin.Context, request IRequest) (bool, *customError.CustomError) {
	if err := request.ValidateRequest(ctx); err != nil {
		return false, err
	}

	if err := request.ValidateAuth(ctx); err != nil {
		return false, err
	}

	return true, nil
}

var Validator *ReqValidator

func init() {
	Validator = &ReqValidator{}
}
