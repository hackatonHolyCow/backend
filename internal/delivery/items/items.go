package items

import (
	"hackathon/backend/internal/service"
	"hackathon/backend/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemsDelivery struct {
	service *service.Service
}

func New(svc *service.Service, router *gin.RouterGroup) {
	delivery := ItemsDelivery{
		service: svc,
	}

	route := router.Group("/items")
	delivery.List(route)
}

func (i *ItemsDelivery) List(route *gin.RouterGroup) {
	route.GET("/", func(c *gin.Context) {
		response, err := i.service.Items.List(c)
		if err != nil {
			c.JSON(errors.HTTPCode(err), gin.H{
				"err,or": gin.H{
					"code":  errors.HTTPCode(err),
					"error": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, response)
	})
}
