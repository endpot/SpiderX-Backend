package tracker

import (
	"database/sql"
	"fmt"
	trackerReq "github.com/endpot/SpiderX-Backend/app/controller/request/tracker"
	"github.com/endpot/SpiderX-Backend/app/domain/model"
	"github.com/endpot/SpiderX-Backend/app/infra/db"
	"github.com/endpot/SpiderX-Backend/app/infra/util/convert"
	customError "github.com/endpot/SpiderX-Backend/app/infra/util/error"
	trackerUtil "github.com/endpot/SpiderX-Backend/app/infra/util/tracker"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net"
	"time"
)

// 生成 Scrape 结果
func GenScrapeResult(ctx *gin.Context, req *trackerReq.ScrapeRequest) (model.TorrentSlice, *customError.CustomError) {
	var infoHashSlice []string

	for _, infoHash := range req.InfoHashSlice {
		infoHashSlice = append(infoHashSlice, trackerUtil.RestoreToHexString(infoHash))
	}

	torrentSlice, err := model.Torrents(model.TorrentWhere.InfoHash.IN(infoHashSlice)).All(ctx, db.DB)
	if err != nil {
		return nil, customError.NewBadRequestError("TRACKER__INVALID_TORRENT")
	}

	return torrentSlice, nil
}

// 处理客户端上报的请求
func DealWithClientReport(ctx *gin.Context, req *trackerReq.AnnounceRequest) (*model.Torrent, model.PeerSlice, *customError.CustomError) {
	// 查询种子
	torrent, err := model.Torrents(model.TorrentWhere.InfoHash.EQ(req.InfoHash)).One(ctx, db.DB)
	if err != nil || torrent == nil {
		return nil, nil, customError.NewBadRequestError("TRACKER__INVALID_TORRENT")
	}

	// 更新数据
	if err := updateData(ctx, torrent, req); err != nil {
		return nil, nil, customError.NewBadRequestError("TRACKER__INVALID_TORRENT")
	}

	// 查找同伴列表
	peerSlice, err := retrievePeerList(ctx, torrent, req)
	if err != nil {
		return nil, nil, customError.NewBadRequestError("TRACKER__INVALID_PARAMS")
	}

	return torrent, peerSlice, nil
}

// 更新数据
func updateData(ctx *gin.Context, torrent *model.Torrent, req *trackerReq.AnnounceRequest) error {
	exists, err := model.Peers(
		model.PeerWhere.TorrentID.EQ(torrent.ID),
		model.PeerWhere.PeerID.EQ(req.PeerID),
	).Exists(ctx, db.DB)
	if err != nil {
		return err
	}

	if exists {
		if req.Event == "stopped" {
			if err := handleStoppedEvent(ctx, torrent, req); err != nil {
				return err
			}
		} else {
			if err := handleNormalEvent(ctx, torrent, req); err != nil {
				return err
			}
		}
	} else {
		if err := handleNewPeer(ctx, torrent, req); err != nil {
			return err
		}
	}

	return nil
}

// handle stopped
func handleStoppedEvent(ctx *gin.Context, torrent *model.Torrent, req *trackerReq.AnnounceRequest) error {
	currentPeer, err := model.Peers(
		model.PeerWhere.TorrentID.EQ(torrent.ID),
		model.PeerWhere.PeerID.EQ(req.PeerID),
	).One(ctx, db.DB)
	if err != nil {
		return err
	}

	// TODO: UserID
	snatcher, err := model.Snatchers(
		model.SnatcherWhere.TorrentID.EQ(torrent.ID),
		model.SnatcherWhere.UserID.EQ(0),
	).One(ctx, db.DB)
	if err != nil {
		return nil
	}

	snatcher.UploadedBytes = uint64(req.UploadedBytes)
	snatcher.DownloadedBytes = uint64(req.DownloadedBytes)
	snatcher.LeftBytes = uint64(req.LeftBytes)

	announceTime := uint(time.Now().Unix() - currentPeer.UpdatedAt.Time.Unix())
	if convert.ParseUint8ToBool(currentPeer.IsSeeder) {
		snatcher.SeedTime = snatcher.SeedTime + announceTime
	} else {
		snatcher.LeechTime = snatcher.LeechTime + announceTime
	}

	// 更新下载者数据
	if _, err := snatcher.Update(ctx, db.DB, boil.Infer()); err != nil {
		return err
	}

	// 删除做种者
	if _, err := currentPeer.Delete(ctx, db.DB); err != nil {
		return err
	}

	return nil
}

// handle normal event
func handleNormalEvent(ctx *gin.Context, torrent *model.Torrent, req *trackerReq.AnnounceRequest) error {
	currentPeer, err := model.Peers(
		model.PeerWhere.TorrentID.EQ(torrent.ID),
		model.PeerWhere.PeerID.EQ(req.PeerID),
	).One(ctx, db.DB)
	if err != nil {
		return err
	}

	currentPeer.Ipv4 = req.IPv4
	currentPeer.Ipv6 = req.IPv6
	currentPeer.UploadedBytes = uint64(req.UploadedBytes)
	currentPeer.DownloadedBytes = uint64(req.DownloadedBytes)
	currentPeer.LeftBytes = uint64(req.LeftBytes)
	currentPeer.IsSeeder = convert.ParseBoolToUint8(req.LeftBytes == 0)
	currentPeer.Agent = req.Agent
	if req.Event == "completed" {
		currentPeer.FinishedAt = null.TimeFrom(time.Now())
	}
	if _, err := currentPeer.Update(ctx, db.DB, boil.Infer()); err != nil {
		return err
	}

	// TODO: UserID
	snatcher, err := model.Snatchers(
		model.SnatcherWhere.TorrentID.EQ(torrent.ID),
		model.SnatcherWhere.UserID.EQ(0),
	).One(ctx, db.DB)
	if err != nil {
		return nil
	}

	snatcher.UploadedBytes = uint64(req.UploadedBytes)
	snatcher.DownloadedBytes = uint64(req.DownloadedBytes)
	snatcher.LeftBytes = uint64(req.LeftBytes)

	announceTime := uint(time.Now().Unix() - currentPeer.UpdatedAt.Time.Unix())
	if convert.ParseUint8ToBool(currentPeer.IsSeeder) {
		snatcher.SeedTime = snatcher.SeedTime + announceTime
	} else {
		snatcher.LeechTime = snatcher.LeechTime + announceTime
	}

	// 更新下载者数据
	if _, err := snatcher.Update(ctx, db.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// handle new peer
func handleNewPeer(ctx *gin.Context, torrent *model.Torrent, req *trackerReq.AnnounceRequest) error {
	currentPeer := &model.Peer{
		TorrentID:             torrent.ID,
		PeerID:                req.PeerID,
		Passkey:               req.Passkey,
		Ipv4:                  req.IPv4,
		Ipv6:                  req.IPv6,
		Port:                  req.Port,
		UploadedBytes:         uint64(req.UploadedBytes),
		DownloadedBytes:       uint64(req.DownloadedBytes),
		OffsetUploadedBytes:   uint64(req.UploadedBytes),
		OffsetDownloadedBytes: uint64(req.DownloadedBytes),
		LeftBytes:             uint64(req.LeftBytes),
		IsSeeder:              convert.ParseBoolToUint8(req.LeftBytes == 0),
		IsConnectable:         convert.ParseBoolToUint8(checkConnectable(req.IPv4, req.IPv6, req.Port)),
		Agent:                 req.Agent,
	}

	if err := currentPeer.Insert(ctx, db.DB, boil.Infer()); err != nil {
		return err
	}

	// TODO: UserID
	snatcher, err := model.Snatchers(
		model.SnatcherWhere.TorrentID.EQ(torrent.ID),
		model.SnatcherWhere.UserID.EQ(0),
	).One(ctx, db.DB)
	if err != nil && err != sql.ErrNoRows {
		return nil
	}

	var noSnatcher bool

	if snatcher == nil {
		snatcher = &model.Snatcher{}
		noSnatcher = true
	} else {
		noSnatcher = false
	}

	snatcher.UploadedBytes = uint64(req.UploadedBytes)
	snatcher.DownloadedBytes = uint64(req.DownloadedBytes)
	snatcher.LeftBytes = uint64(req.LeftBytes)

	announceTime := uint(time.Now().Unix() - currentPeer.UpdatedAt.Time.Unix())
	if convert.ParseUint8ToBool(currentPeer.IsSeeder) {
		snatcher.SeedTime = snatcher.SeedTime + announceTime
	} else {
		snatcher.LeechTime = snatcher.LeechTime + announceTime
	}

	if noSnatcher {
		snatcher.Insert(ctx, db.DB, boil.Infer())
	} else {
		// 更新下载者数据
		if _, err := snatcher.Update(ctx, db.DB, boil.Infer()); err != nil {
			return err
		}
	}

	return nil
}

// 检查连通性
func checkConnectable(ipv4 string, ipv6 string, port uint16) bool {
	connectable := false

	if ipv6 != "" {
		if _, err := net.DialTimeout(
			"tcp6",
			fmt.Sprintf("[%v]:%v", ipv6, port),
			5*time.Second,
		); err == nil {
			connectable = true
		}
	}

	if !connectable && ipv4 != "" {
		if _, err := net.DialTimeout(
			"tcp4",
			fmt.Sprintf("%v:%v", ipv4, port),
			5*time.Second,
		); err == nil {
			connectable = true
		}
	}

	return connectable
}

// 获取同伴列表
// 如果当前用户为做种者，只返回其它非做种者
func retrievePeerList(ctx *gin.Context, torrent *model.Torrent, req *trackerReq.AnnounceRequest) (model.PeerSlice, error) {
	columns := []string{"ipv4", "ipv6", "port"}
	if req.Compact != 1 || req.NoPeerID != 1 {
		columns = append(columns, "peer_id")
	}

	isSeeder := req.LeftBytes == 0
	if isSeeder {
		return model.Peers(
			qm.Select(columns...),
			model.PeerWhere.TorrentID.EQ(torrent.ID),
			model.PeerWhere.PeerID.NEQ(req.PeerID),
			model.PeerWhere.IsSeeder.NEQ(1),
			qm.OrderBy("RAND()"),
			qm.Limit(calPeerLimitCount(req.NumWanted)),
		).All(ctx, db.DB)
	}

	return model.Peers(
		qm.Select(columns...),
		model.PeerWhere.TorrentID.EQ(torrent.ID),
		model.PeerWhere.PeerID.NEQ(req.PeerID),
		qm.OrderBy("RAND()"),
		qm.Limit(calPeerLimitCount(req.NumWanted)),
	).All(ctx, db.DB)
}

// 获取同伴数量
func calPeerLimitCount(numWanted uint8) int {
	return int(numWanted)
}
