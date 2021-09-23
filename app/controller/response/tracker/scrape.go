package tracker

import (
	"github.com/anacrolix/torrent/bencode"
	"github.com/endpot/SpiderX-Backend/app/domain/model"
	trackerUtil "github.com/endpot/SpiderX-Backend/app/infra/util/tracker"
)

type ScrapeItem struct {
	SeederCount   uint `bencode:"complete"`   // 当前做种数量
	SnatcherCount uint `bencode:"downloaded"` // 已完成下载总数
	LeecherCount  uint `bencode:"incomplete"` // 正在下载数量
}

type ScrapeResult struct {
	Files map[string]*ScrapeItem `bencode:"files"`
}

type ScrapeResponse struct {
}

func (s ScrapeResponse) Serialize(singleModel interface{}) interface{} {
	return nil
}

func (s ScrapeResponse) Paginate(modelSlice interface{}) interface{} {
	return nil
}

func (s ScrapeResponse) Bencode(modelSlice interface{}) string {
	torrentSlice, ok := modelSlice.(model.TorrentSlice)
	if !ok {
		return string(bencode.MustMarshal(map[string]string{
			"failure reason": "Bad Torrent",
		}))
	}

	scrapeResult := &ScrapeResult{
		Files: make(map[string]*ScrapeItem),
	}
	for _, torrent := range torrentSlice {
		scrapeResult.Files[trackerUtil.RestoreToByteString(torrent.InfoHash)] = &ScrapeItem{
			SeederCount:   torrent.SeederCount,
			SnatcherCount: torrent.SnatcherCount,
			LeecherCount:  torrent.LeecherCount,
		}
	}

	return string(bencode.MustMarshal(scrapeResult))
}
