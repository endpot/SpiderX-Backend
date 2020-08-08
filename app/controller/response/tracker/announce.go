package tracker

import (
	"bytes"
	"encoding/binary"
	"github.com/anacrolix/torrent/bencode"
	trackerReq "github.com/endpot/SpiderX-Backend/app/controller/request/tracker"
	"github.com/endpot/SpiderX-Backend/app/domain/model"
	trackerUtil "github.com/endpot/SpiderX-Backend/app/infra/util/tracker"
	"github.com/volatiletech/null/v8"
	"math"
	"net"
	"time"
)

type PeerItem struct {
	IP     string `bencode:"ip"`
	Port   uint   `bencode:"port"`
	PeerID string `bencode:"peer id,omitempty"`
}

type AnnounceResult struct {
	Interval      uint        `bencode:"interval"`   // 请求时间间隔
	SeederCount   uint        `bencode:"complete"`   // 当前做种数量
	SnatcherCount uint        `bencode:"downloaded"` // 已完成下载总数
	LeecherCount  uint        `bencode:"incomplete"` // 正在下载数量
	Peers         []*PeerItem `bencode:"peers"`      // 同伴
}

type AnnounceCompactResult struct {
	Interval      uint   `bencode:"interval"`   // 请求时间间隔
	SeederCount   uint   `bencode:"complete"`   // 当前做种数量
	SnatcherCount uint   `bencode:"downloaded"` // 已完成下载总数
	LeecherCount  uint   `bencode:"incomplete"` // 正在下载数量
	Peers         string `bencode:"peers"`      // IPv4 同伴
	Peers6        string `bencode:"peers6"`     // IPv6 同伴
}

type AnnounceResponse struct {
}

func (s AnnounceResponse) Serialize(_ interface{}) interface{} {
	return nil
}

func (s AnnounceResponse) Paginate(_ interface{}) interface{} {
	return nil
}

func (s AnnounceResponse) Bencode(torrent *model.Torrent, modelSlice interface{}, req *trackerReq.AnnounceRequest) string {
	peerSlice, ok := modelSlice.(model.PeerSlice)
	if !ok {
		return string(bencode.MustMarshal(map[string]string{
			"failure reason": "Bad Torrent",
		}))
	}

	if req.Compact == 1 {
		peers, peers6 := genCompactPeers(peerSlice)
		compactResult := &AnnounceCompactResult{
			Interval:      genInterval(torrent.CreatedAt),
			SeederCount:   torrent.SeederCount,
			SnatcherCount: torrent.SnatcherCount,
			LeecherCount:  torrent.LeecherCount,
			Peers:         peers,
			Peers6:        peers6,
		}
		return string(bencode.MustMarshal(compactResult))
	}

	result := &AnnounceResult{
		Interval:      genInterval(torrent.CreatedAt),
		SeederCount:   torrent.SeederCount,
		SnatcherCount: torrent.SnatcherCount,
		LeecherCount:  torrent.LeecherCount,
		Peers:         genPeers(peerSlice),
	}
	return string(bencode.MustMarshal(result))
}

// 计算请求时间间隔，防止大量请求同时发生
// interval = c/(1+a*e^(-kx)); c = 3600; a = 5; k = 0.0001
// 新种子 10 分钟，旧种子 60 分钟
func genInterval(createdAt null.Time) uint {
	diffSeconds := time.Now().Unix() - createdAt.Time.Unix()
	if diffSeconds <= 0 {
		return 0
	}
	return uint(math.Round(3600 / (1 + 5*math.Exp(float64(-diffSeconds)/60/10000))))
}

// 生成非压缩的同伴列表
func genPeers(peerSlice model.PeerSlice) []*PeerItem {
	var peers []*PeerItem

	for _, peer := range peerSlice {
		if peer.Ipv4 != "" {
			peers = append(peers, &PeerItem{
				IP:     peer.Ipv4,
				Port:   uint(peer.Port),
				PeerID: trackerUtil.RestoreToByteString(peer.PeerID),
			})
		}

		if peer.Ipv6 != "" {
			peers = append(peers, &PeerItem{
				IP:     peer.Ipv6,
				Port:   uint(peer.Port),
				PeerID: trackerUtil.RestoreToByteString(peer.PeerID),
			})
		}
	}

	return peers
}

// 生成压缩的同伴列表
func genCompactPeers(peerSlice model.PeerSlice) (string, string) {
	var peers []byte
	var peers6 []byte

	buf := new(bytes.Buffer)
	byteOrder := binary.BigEndian

	for _, peer := range peerSlice {
		if ip := net.ParseIP(peer.Ipv4).To4(); ip != nil {
			if buf.Reset(); binary.Write(buf, byteOrder, ip) == nil && binary.Write(buf, byteOrder, uint16(peer.Port)) == nil {
				peers = append(peers, buf.Bytes()...)
			}
		}

		if ip := net.ParseIP(peer.Ipv6).To16(); ip != nil {
			if buf.Reset(); binary.Write(buf, byteOrder, ip) == nil && binary.Write(buf, byteOrder, uint16(peer.Port)) == nil {
				peers6 = append(peers6, buf.Bytes()...)
			}
		}
	}

	return string(peers), string(peers6)
}
