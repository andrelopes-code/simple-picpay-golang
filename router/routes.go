package router

import (
	"github.com/andrelopes-code/simple-picpay-golang/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/transaction", handler.CreateTransaction)
	}
}
