package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/config"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := config.GetString("global.indexHtml")
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			utils.Logger().Errorf("readfile failed. %v", err)
			utils.GinJson(c, nil, err)
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", bytes)
	}
}
