package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"detail": msg,
		"code":   code,
	})
}

func sendSuccess(ctx *gin.Context, code int, op string, data any) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"detail": fmt.Sprintf("operation: %s concluded with success", op),
		"data":   data,
	})
}
