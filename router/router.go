package router

import (
	"github.com/andrelopes-code/simple-picpay-golang/cfg"
	"github.com/andrelopes-code/simple-picpay-golang/handler"
	"github.com/gin-gonic/gin"
)

var log = cfg.GetLogger("router")

func Init() {
	log.Info("Initializing router...")
	// Initialize handler package
	handler.Init()
	// Create router and initialize routes
	router := gin.Default()
	initializeRoutes(router)
	router.Run(": 8080")
}
