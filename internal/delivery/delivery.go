package delivery

import (
	"hackathon/backend/internal/delivery/items"
	"hackathon/backend/internal/delivery/orders"
	"hackathon/backend/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(svc *service.Service) *gin.Engine {
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Requested-With", "X-HTTP-Method-Override"},
	}))

	apiRouter := app.Group("/api/v1")
	orders.New(svc, apiRouter)
	items.New(svc, apiRouter)
	return app
}
