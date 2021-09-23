package tracker

import (
	"github.com/anacrolix/torrent/bencode"
	"github.com/endpot/SpiderX-Backend/app/controller/request"
	"github.com/endpot/SpiderX-Backend/app/controller/request/tracker"
	trackerResp "github.com/endpot/SpiderX-Backend/app/controller/response/tracker"
	trackerService "github.com/endpot/SpiderX-Backend/app/domain/service/tracker"
	"github.com/endpot/SpiderX-Backend/app/infra/util/http"
	"github.com/gin-gonic/gin"
	"log"
)

func Announce(ctx *gin.Context) {
	// 使用 AnnounceRequest 校验请求
	req := &tracker.AnnounceRequest{}
	if _, err := request.Validator.Validate(
		ctx,
		req,
	); err != nil {
		ctx.String(err.Status, packFailReason(err.Error()))
		return
	}

	torrent, peerSlice, err := trackerService.DealWithClientReport(ctx, req)
	if err != nil {
		log.Println(err)
		ctx.String(err.Status, packFailReason(err.Error()))
		return
	}
	resp := &trackerResp.AnnounceResponse{}
	ctx.String(http.StatusOK, resp.Bencode(torrent, peerSlice, req))
}

func Scrape(ctx *gin.Context) {
	// 使用 ScrapeRequest 校验请求
	req := &tracker.ScrapeRequest{}
	if _, err := request.Validator.Validate(
		ctx,
		req,
	); err != nil {
		ctx.String(err.Status, packFailReason(err.Error()))
		return
	}

	// 查找符合条件的种子
	torrentSlice, _ := trackerService.GenScrapeResult(ctx, req)
	resp := &trackerResp.ScrapeResponse{}
	ctx.String(http.StatusOK, resp.Bencode(torrentSlice))
}

// Bencode 错误原因
func packFailReason(reason string) string {
	return string(bencode.MustMarshal(map[string]string{
		"failure reason": reason,
	}))
}
