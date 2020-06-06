package cgi

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/internal/service"
	"github.com/day-dreams/vshare.zhangnan.xyz/internal/service/hls"
	"github.com/day-dreams/vshare.zhangnan.xyz/internal/utils"
)

func VideoInfoDemo() gin.HandlerFunc {
	return func(c *gin.Context) {
		info, err := service.VideoInfoGet(c, "/data/vmoon/detachment.mkv")
		utils.GinJson(c, info, err)
	}
}
func M3u8PlayList() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &hls.ReqM3u8PlayList{
			Vid:  c.Query("vid"),
			Path: "/api/video/hls/segment",
		}

		//loader := hls.LiveFfmpegLoader{}
		loader := hls.DiskFfmpegLoader{}

		res, err := loader.M3u8PlayList(c, req)
		if err != nil {
			utils.GinJson(c, nil, err)
			return
		}

		c.Writer.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
		_, _ = c.Writer.Write([]byte(res.PlayListContent))
	}
}

func M3u8Segment() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &hls.ReqM3u8Segment{
			Vid:   c.Query("vid"),
			Index: 0,
		}
		index, err := strconv.Atoi(c.Query("segment"))
		if err != nil {
			utils.GinJson(c, nil, err)
			return
		}
		req.Index = index

		//loader := hls.LiveFfmpegLoader{}
		loader := hls.DiskFfmpegLoader{}

		res, err := loader.M3u8Segment(c, req)
		if err != nil {
			utils.GinJson(c, nil, err)
			return
		}

		c.Writer.Header().Set("Content-Type", "video/mp2t")
		_, _ = c.Writer.Write(res.Content)
	}
}
