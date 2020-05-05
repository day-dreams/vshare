package handler

import (
	"strconv"

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

func M3u8Segment() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.Logger().Debugf("......")
		req := &hls.ReqM3u8Segment{
			Vid:   c.Query("vid"),
			Index: 0,
		}
		utils.Logger().Debugf("......")
		index, err := strconv.Atoi(c.Query("segment"))
		utils.Logger().Debugf("......")
		if err != nil {
			utils.GinJson(c, nil, err)
			return
		}
		req.Index = index

		res, err := hls.M3u8Segment(c, req)
		if err != nil {
			utils.GinJson(c, nil, err)
			return
		}

		c.Writer.Header().Set("Content-Type", "video/mp2t")
		_, _ = c.Writer.Write(res.Content)
	}
}
