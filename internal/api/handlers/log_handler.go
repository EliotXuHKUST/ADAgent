package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Event struct {
	GameID    string                 `json:"game_id" binding:"required"`
	UserID    string                 `json:"user_id" binding:"required"`
	Timestamp int64                  `json:"timestamp" binding:"required"`
	EventName string                 `json:"event_name" binding:"required"`
	Context   map[string]interface{} `json:"context"`
}

type LogHandler struct {
	logger *zap.Logger
}

func NewLogHandler(logger *zap.Logger) *LogHandler {
	return &LogHandler{logger: logger}
}

func (h *LogHandler) CollectEvent(c *gin.Context) {
	var event Event
	if err := c.ShouldBindJSON(&event); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Push to Kafka/Redpanda
	h.logger.Info("Received event",
		zap.String("game_id", event.GameID),
		zap.String("event", event.EventName),
		zap.String("user_id", event.UserID),
	)

	// Mock response strategy for MVP Phase 1
	c.JSON(http.StatusOK, gin.H{
		"status": "received",
		"trace_id": "mock-trace-id",
	})
}
