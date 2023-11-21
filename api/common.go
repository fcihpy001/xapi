package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func response(ctx *gin.Context, status int, code int, msg string, data interface{}) {
	ctx.AbortWithStatusJSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ErrorResp(ctx *gin.Context, code int, msg string) {
	response(ctx, http.StatusOK, code, msg, nil)
}

func SuccessResp(ctx *gin.Context, msg string, data interface{}) {
	message := msg
	if len(message) == 0 {
		message = "success"
	}
	response(ctx, http.StatusOK, 200, message, data)
}
