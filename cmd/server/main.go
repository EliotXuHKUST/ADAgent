package main

import (
	"github.com/EliotXuHKUST/ADAgent/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Initialize Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize Handlers
	logHandler := handlers.NewLogHandler(logger)

	// Setup Router
	r := gin.Default()
	
	v1 := r.Group("/v1")
	{
		v1.POST("/events/collect", logHandler.CollectEvent)
	}

	logger.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
