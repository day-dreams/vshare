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
		v := video{
			VID:   vid,
			Title: vtitle,
			Desc:  vdesc,
			TxCloud: struct {
				AppId  string
				FileId string
			}{AppId: appid, FileId: fileid},
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
