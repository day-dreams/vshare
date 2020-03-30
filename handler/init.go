package handler

import (
	"fmt"
	"time"

	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func init() {
	setupVideo()
	setupRoom()
	go func() {

		t := time.NewTicker(time.Second * 10)
		for {
			select {
			case _ = <-t.C:
				inspect()
			}
		}
	}()
}

func inspect() {
	mutex.RLock()
	defer mutex.RUnlock()
	fmt.Printf("rooms  :\t%s\n", utils.ToPrettyJSON(rid2room))
	fmt.Printf("clients:\t%s\n", utils.ToPrettyJSON(cid2client))
}
