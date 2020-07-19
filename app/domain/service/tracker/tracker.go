package tracker

import (
	trackerReq "github.com/endpot/SpiderX-Backend/app/controller/request/tracker"
	"github.com/endpot/SpiderX-Backend/app/domain/model"
	"github.com/endpot/SpiderX-Backend/app/infra/db"
	customError "github.com/endpot/SpiderX-Backend/app/infra/util/error"
	trackerUtil "github.com/endpot/SpiderX-Backend/app/infra/util/tracker"
	"github.com/gin-gonic/gin"
)

// 生成 Scrape 结果
func GenScrapeResult(ctx *gin.Context, req *trackerReq.ScrapeRequest) (model.TorrentSlice, *customError.CustomError) {
	var infoHashSlice []string

	for _, infoHash := range req.InfoHashSlice {
		infoHashSlice = append(infoHashSlice, trackerUtil.RestoreToHexString(infoHash))
	}

	torrentSlice, err := model.Torrents(model.TorrentWhere.InfoHash.IN(infoHashSlice)).All(ctx, db.DB)
	if err != nil {
		return nil, customError.NewBadRequestError("")
	}

	return torrentSlice, nil
}
