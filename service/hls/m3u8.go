package hls

import (
	"context"
	"fmt"

	"github.com/day-dreams/vshare.zhangnan.xyz/config"
	"github.com/day-dreams/vshare.zhangnan.xyz/service"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

type ReqM3u8PlayList struct {
	Vid  string
	Path string
}

type ResM3u8PlayList struct {
	PlayListContent string
}

func M3u8PlayList(ctx context.Context, req *ReqM3u8PlayList) (*ResM3u8PlayList, error) {

	key := fmt.Sprintf("videos.%s.path", req.Vid)
	path := config.GetString(key)
	vInfo, err := service.VideoInfoGet(ctx, path)
	if err != nil {
		return nil, err
	}

	res := &ResM3u8PlayList{PlayListContent: ""}
	total := vInfo.Duration
	DPB := config.GetInt("global.durationPerBatch")
	cur := 0

	// header + play list
	res.PlayListContent += fmt.Sprintf("#EXTM3U\n#EXT-X-TARGETDURATION:%d\n", DPB)
	for ; float64(cur*DPB) <= total; cur += 1 {
		utils.Logger().Debugf("play list:%v,%v,%v,%v", cur, float64(cur*DPB), total, float64(cur*DPB) <= total)
		duration := float64(DPB)
		if float64(cur*DPB)+duration > total {
			duration = total - float64(cur*DPB)
		}
		// [cur*DPB,cur*DPB+duration)
		res.PlayListContent += fmt.Sprintf("#EXTINF:%.2f,\n%s?segment=%d.ts\n", duration, req.Path, cur)
	}
	return res, nil
}
