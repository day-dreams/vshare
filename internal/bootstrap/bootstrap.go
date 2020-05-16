package bootstrap

import (
	"github.com/day-dreams/vshare.zhangnan.xyz/internal/config"
	"github.com/day-dreams/vshare.zhangnan.xyz/internal/utils"
)

func Bootstrap() {
	config.InitConfig()
	utils.InitLogger()
}
