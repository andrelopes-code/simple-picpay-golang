package router

import (
	"github.com/andrelopes-code/simple-picpay-golang/docs"
	"github.com/andrelopes-code/simple-picpay-golang/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		// Routes for transactions
		v1.POST("/transaction", handler.CreateTransaction)
	}
	// Route for swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler()))
}
