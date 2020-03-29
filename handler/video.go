package handler

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"

	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

type video struct {
	Title   string
	Desc    string
	TxCloud struct {
		AppId  string
		FileId string
	}
}

var (
	videos []video
)

func init() {
	vfile := utils.VFile
	bytes, err := ioutil.ReadFile(vfile)
	if err != nil {
		panic(err)
	}
	data := string(bytes)

	gjson.Get(data, "videos").ForEach(func(key, value gjson.Result) bool {
		vtitle := value.Get("vtitle").String()
		vdesc := value.Get("vdesc").String()
		fileid := value.Get("tcloud").Get("fileid").String()
		appid := value.Get("tcloud").Get("appid").String()
		videos = append(videos, video{
			Title: vtitle,
			Desc:  vdesc,
			TxCloud: struct {
				AppId  string
				FileId string
			}{AppId: appid, FileId: fileid},
		})
		return true
	})
}
func VideoList() gin.HandlerFunc {

	return func(c *gin.Context) {
		utils.GinJson(c, videos, nil)
	}
}
