package tracker

import (
	"github.com/endpot/SpiderX-Backend/app/controller/request"
	"github.com/endpot/SpiderX-Backend/app/controller/request/tracker"
	trackerResp "github.com/endpot/SpiderX-Backend/app/controller/response/tracker"
	trackerService "github.com/endpot/SpiderX-Backend/app/domain/service/tracker"
	"github.com/endpot/SpiderX-Backend/app/infra/util/http"
	"github.com/gin-gonic/gin"
)

func Announce(ctx *gin.Context) {
	//
}

func Scrape(ctx *gin.Context) {
	// 使用 ScrapeRequest 校验请求
	req := &tracker.ScrapeRequest{}
	if _, err := request.Validator.Validate(
		ctx,
		req,
	); err != nil {
		ctx.JSON(err.Status, err.Serialize())
		return
	}

	// 查找符合条件的种子
	torrentSlice, _ := trackerService.GenScrapeResult(ctx, req)
	resp := &trackerResp.ScrapeResponse{}
	ctx.String(http.StatusOK, resp.Bencode(torrentSlice))
}
