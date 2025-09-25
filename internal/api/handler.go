package api

import (
	"ApplicationManager/internal/logging"
	"ApplicationManager/internal/metrics"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateRequest(c *gin.Context) {
	logging.Logger.Info("Received create request")
	metrics.RequestCount.Inc()
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
