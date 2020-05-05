package bootstrap

import (
	"github.com/day-dreams/vshare.zhangnan.xyz/config"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func Bootstrap() {
	config.InitConfig()
	utils.InitLogger()
}
