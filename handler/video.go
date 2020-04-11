package handler

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"

	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

type video struct {
	VID     string
	Title   string
	Desc    string
	TxCloud struct {
		AppId  string
		FileId string
	}
	Hls struct {
		M3u8  string
		TsDir string
	}
}

var (
	videos    []video
	vid2video = map[string]video{}
)

func setupVideo() {
	vfile := utils.VFile
	bytes, err := ioutil.ReadFile(vfile)
	if err != nil {
		panic(err)
	}
	data := string(bytes)

	gjson.Get(data, "videos").ForEach(func(key, value gjson.Result) bool {
		vid := value.Get("vid").String()
		vtitle := value.Get("vtitle").String()
		vdesc := value.Get("vdesc").String()
		fileid := value.Get("tcloud").Get("fileid").String()
		appid := value.Get("tcloud").Get("appid").String()
		m3u8 := value.Get("hls").Get("m3u8").String()
		tsdir := value.Get("hls").Get("tsdir").String()
		v := video{
			VID:   vid,
			Title: vtitle,
			Desc:  vdesc,
			TxCloud: struct {
				AppId  string
				FileId string
			}{AppId: appid, FileId: fileid},
			Hls: struct {
				M3u8  string
				TsDir string
			}{M3u8: m3u8, TsDir: tsdir},
		}
		videos = append(videos, v)
		vid2video[vid] = v
		return true
	})
}
func VideoList() gin.HandlerFunc {

	return func(c *gin.Context) {
		utils.GinJson(c, videos, nil)
	}
}
