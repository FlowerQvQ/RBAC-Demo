package wapper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func ResSuccess(c *gin.Context, v interface{}) {
	result := gin.H{}
	result["code"] = 200
	result["msg"] = "success"
	if v == nil || reflect.ValueOf(v).IsZero() {
		result["data"] = struct{}{}
	} else {
		result["data"] = v
	}

	c.JSON(http.StatusOK, result)
}

func ResError(c *gin.Context, errCode ErrorCode) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    errCode.Code,
		"message": errCode.Message,
	})
}
