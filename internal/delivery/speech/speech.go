package speech

import (
	"hackathon/backend/entity"
	"hackathon/backend/internal/service"
	"hackathon/backend/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SpeechDelivery struct {
	service *service.Service
}

func New(svc *service.Service, router *gin.RouterGroup) {
	delivery := &SpeechDelivery{
		service: svc,
	}

	route := router.Group("/speech")
	delivery.TextToSpeach(route)
}

func (s *SpeechDelivery) TextToSpeach(route *gin.RouterGroup) {
	route.POST("", func(c *gin.Context) {
		var request entity.TextToSpeechRequest
		if err := c.Bind(&request); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": gin.H{
					"code":  http.StatusUnprocessableEntity,
					"error": err.Error(),
				},
			})

			return
		}

		if request.Text == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{
					"code":  http.StatusBadRequest,
					"error": "text is required",
				},
			})

			return
		}

		response, err := s.service.Speech.TextToSpeach(c, request.Text)
		if err != nil {
			c.JSON(errors.HTTPCode(err), gin.H{
				"error": gin.H{
					"code":  errors.HTTPCode(err),
					"error": err.Error(),
				},
			})

			return
		}

		c.Data(http.StatusOK, "audio/mp3", response)
	})
}
