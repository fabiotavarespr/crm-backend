package controllers

import (
	"net/http"

	"github.com/fabiotavarespr/crm-backend/services"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthService services.HealthService
}

func NewHealthController(healthService services.HealthService) HealthController {
	return HealthController{healthService}
}

func (hc *HealthController) CheckHealth(ctx *gin.Context) {
	newHealth, err := hc.healthService.CheckHealth()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newHealth)
}
