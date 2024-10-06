package service

import (
	"hackathon/backend/internal/repository"
	"hackathon/backend/internal/service/items"
	"hackathon/backend/internal/service/orders"
	"hackathon/backend/internal/service/speech"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/openai/openai-go"
)

type Service struct {
	Orders orders.OrdersService
	Items  items.ItemsService
	Speech speech.SpeechService
}

func New(repo *repository.Repository, mpConfig config.Config, openAIClient *openai.Client) *Service {
	return &Service{
		Orders: orders.New(repo, mpConfig),
		Items:  items.New(repo),
		Speech: speech.New(openAIClient),
	}
}
