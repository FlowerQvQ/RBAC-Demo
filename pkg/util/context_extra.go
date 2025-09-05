package util

import "github.com/gin-gonic/gin"

const (
	requestIDKey = "requestID"
)

func GetRequestID(context *gin.Context) string {
	return context.GetString(requestIDKey)
}
