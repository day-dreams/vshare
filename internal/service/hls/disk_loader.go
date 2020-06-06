package hls

import (
	"context"
	"fmt"
	"github.com/day-dreams/vshare.zhangnan.xyz/internal/service"
)

type DiskFfmpegLoader struct {
}

func (d DiskFfmpegLoader) M3u8PlayList(ctx context.Context, req *ReqM3u8PlayList) (*ResM3u8PlayList, error) {
	panic("implement me")
}

func (l *LiveFfmpegLoader) M3u8PlayList(ctx context.Context, req *ReqM3u8PlayList) (*ResM3u8PlayList, error) {
	path := GetVidPath(req.Vid)
	vInfo, err := service.VideoInfoGet(ctx, path)
	if err != nil {
		return nil, err
	}

	res := &ResM3u8PlayList{PlayListContent: ""}
	total := vInfo.Duration
	DPB := GetDPB()
	cur := 0

	// header + play list
	res.PlayListContent += fmt.Sprintf("#EXTM3U\n#EXT-X-TARGETDURATION:%d\n\n", DPB)
	for ; float64(cur*DPB) <= total; cur += 1 {
		duration := float64(DPB)
		if float64(cur*DPB)+duration > total {
			duration = total - float64(cur*DPB)
		}
		// [cur*GetDPB,cur*GetDPB+duration)
		res.PlayListContent += fmt.Sprintf("#EXTINF:%.2f,\n%s?vid=%s&segment=%d\n", duration, req.Path, req.Vid, cur)
	}
	res.PlayListContent += fmt.Sprintf("\n#EXT-X-ENDLIST")
	return res, nil
}
