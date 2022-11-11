package routes

import (
	"github.com/fabiotavarespr/crm-backend/controllers"
	"github.com/gin-gonic/gin"
)

type HealthRouteController struct {
	healthController controllers.HealthController
}

func NewHealthRouteController(healthController controllers.HealthController) HealthRouteController {
	return HealthRouteController{healthController}
}

func (rc *HealthRouteController) HealthRoute(rg *gin.Engine) {
	rg.GET("/health", rc.healthController.CheckHealth)
}
