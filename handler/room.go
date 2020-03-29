package handler

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"

	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

type room struct {
	RID   string
	VID   string
	Video video
}

var (
	rooms    []room
	rid2room = map[string]room{}
)

func setupRoom() {
	bytes, err := ioutil.ReadFile(utils.VFile)
	if err != nil {
		panic(err)
	}
	data := string(bytes)

	gjson.Get(data, "rooms").ForEach(func(key, value gjson.Result) bool {
		vid := value.Get("vid").String()
		rid := value.Get("rid").String()

		r := room{
			RID:   rid,
			VID:   vid,
			Video: vid2video[vid],
		}
		rooms = append(rooms, r)
		rid2room[rid] = r
		return true
	})
}

func RoomInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("rid2room: [%+v]", rid2room)
		rid := c.Query("rid")
		utils.GinJson(c, rid2room[rid], nil)
	}
}
