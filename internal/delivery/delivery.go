package delivery

import (
	"hackathon/backend/internal/delivery/orders"
	"hackathon/backend/internal/service"

	"github.com/gin-gonic/gin"
)

func New(svc *service.Service) *gin.Engine {
	app := gin.Default()
	apiRouter := app.Group("/api/v1")
	orders.New(svc, apiRouter)
	return app
}
