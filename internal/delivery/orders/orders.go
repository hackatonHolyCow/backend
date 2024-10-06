package orders

import (
	"hackathon/backend/entity"
	"hackathon/backend/internal/service"
	"hackathon/backend/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersDelivery struct {
	service *service.Service
}

func New(svc *service.Service, router *gin.RouterGroup) {
	delivery := &OrdersDelivery{
		service: svc,
	}

	route := router.Group("/orders")
	delivery.Create(route)
	delivery.Get(route)
	delivery.List(route)
	delivery.UpdateStatus(route)
}

func (o *OrdersDelivery) Create(route *gin.RouterGroup) {
	route.POST("", func(c *gin.Context) {
		var request entity.CreateOrderRequest
		if err := c.Bind(&request); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": gin.H{
					"code":  http.StatusUnprocessableEntity,
					"error": err.Error(),
				},
			})

			return
		}

		response, err := o.service.Orders.Create(c, &request)
		if err != nil {
			c.JSON(errors.HTTPCode(err), gin.H{
				"error": gin.H{
					"code":  errors.HTTPCode(err),
					"error": err.Error(),
				},
			})

			return
		}

		c.JSON(http.StatusOK, response)
	})
}

func (o *OrdersDelivery) Get(route *gin.RouterGroup) {
	route.GET("/:id", func(c *gin.Context) {
		response, err := o.service.Orders.Get(c, c.Param("id"))
		if err != nil {
			c.JSON(errors.HTTPCode(err), gin.H{
				"error": gin.H{
					"code":  errors.HTTPCode(err),
					"error": err.Error(),
				},
			})

			return
		}

		c.JSON(http.StatusOK, response)
	})
}

func (o *OrdersDelivery) List(route *gin.RouterGroup) {
	route.GET("", func(c *gin.Context) {
		response, err := o.service.Orders.List(c)
		if err != nil {
			c.JSON(errors.HTTPCode(err), gin.H{
				"error": gin.H{
					"code":  errors.HTTPCode(err),
					"error": err.Error(),
				},
			})

			return
		}

		c.JSON(http.StatusOK, response)
	})
}

func (o *OrdersDelivery) UpdateStatus(route *gin.RouterGroup) {
	route.PATCH("/:id", func(c *gin.Context) {
		var request entity.UpdateStatusRequest
		if err := c.Bind(&request); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": gin.H{
					"code":  http.StatusUnprocessableEntity,
					"error": err.Error(),
				},
			})

			return
		}

		if request.Status == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{
					"code":  http.StatusBadRequest,
					"error": "status is required",
				},
			})

			return
		}

		if err := o.service.Orders.UpdateStatus(c, c.Param("id"), string(request.Status)); err != nil {
			c.JSON(errors.HTTPCode(err), gin.H{
				"error": gin.H{
					"code":  errors.HTTPCode(err),
					"error": err.Error(),
				},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "status updated",
		})
	})
}
