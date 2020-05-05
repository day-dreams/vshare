package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/service/hls"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func M3u8PlayList() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &hls.ReqM3u8PlayList{
			Vid:  c.Query("vid"),
			Path: "/api/video/hls/segment",
		}

		res, err := hls.M3u8PlayList(c, req)
		if err != nil {
			utils.GinJson(c, nil, err)
			return
		}

		c.Writer.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
		_, _ = c.Writer.Write([]byte(res.PlayListContent))
	}
}
