package hls

import (
	"context"
	"fmt"
	"github.com/day-dreams/vshare.zhangnan.xyz/internal/config"
)

type M2u8SegmentLoader interface {
	M2u8Segment(ctx context.Context, req *ReqM3u8Segment) (*ResM3u8Segment, error)
	M2u8PlayList(ctx context.Context, req *ReqM3u8PlayList) (*ResM3u8PlayList, error)
}

type ReqM3u8PlayList struct {
	Vid  string
	Path string
}

type ResM3u8PlayList struct {
	PlayListContent string
}

type ReqM3u8Segment struct {
	Vid   string
	Index int
}
type ResM3u8Segment struct {
	Content []byte
}

func GetDPB() int {
	return config.GetInt("global.durationPerBatch")
}
func GetVidPath(vid string) string {
	key := fmt.Sprintf("videos.%s.path", vid)
	return config.GetString(key)
}

func GetVidM3u8Path(vid string) string {
	key := fmt.Sprintf("videos.%s.m3u8dir", vid)
	return config.GetString(key)
}
