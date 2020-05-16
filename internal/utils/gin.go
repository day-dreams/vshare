package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func GinJson(c *gin.Context, data interface{}, err error) {
	res := map[string]interface{}{
		"code": 0,
		"msg":  "ok",
		"data": data,
	}
	if err != nil {
		s := status.Convert(err)
		res["code"] = s.Code()
		res["emsg"] = s.Message()
	}
	c.JSON(http.StatusOK, res)
}
