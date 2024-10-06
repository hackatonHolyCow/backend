package orders

import (
	"context"
	"fmt"
	"hackathon/backend/entity"
	"hackathon/backend/internal/repository"
	"hackathon/backend/pkg/errors"
	"io"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OrdersService interface {
	Create(ctx context.Context, order *entity.CreateOrderRequest) (*entity.Order, error)
	Get(ctx context.Context, id string) (*entity.Order, error)
	List(ctx context.Context) ([]*entity.Order, error)
	TextToSpeach(ctx context.Context, text string) ([]byte, error)
}

type OrdersServiceImpl struct {
	repository *repository.Repository
	mpConfig   config.Config
}

func New(repo *repository.Repository, mpConfig config.Config) OrdersService {
	return &OrdersServiceImpl{
		repository: repo,
		mpConfig:   mpConfig,
	}
}

func (o *OrdersServiceImpl) Create(ctx context.Context, order *entity.CreateOrderRequest) (*entity.Order, error) {
	paymentClient := payment.NewClient(&o.mpConfig)
	paymentResponse, err := paymentClient.Create(ctx, order.Payment)
	if err != nil {
		return nil, errors.Wrap(err, "orders: OrdersService.Create paymentClient.Create error")
	}
	response, err := o.repository.Orders.Create(ctx, &entity.Order{
		ID:          uuid.NewString(),
		TotalAmount: int(order.Payment.TransactionAmount),
		Table:       order.Table,
		PaymentID:   fmt.Sprint(paymentResponse.ID),
	})

	if err != nil {
		return nil, errors.Wrap(err, "orders: OrdersService.Create repository.Orders.Create error")
	}

	for _, item := range order.Items {
		_, err := o.repository.OrderItems.Create(ctx, &entity.OrderItem{
			OrderID:  response.ID,
			ItemID:   item.ID,
			Quantity: item.Quantity,
			Comments: item.Comments,
		})

		if err != nil {
			return nil, errors.Wrap(err, "orders: OrdersService.Create repository.OrderItems.Create error")
		}
	}

	items, err := o.repository.Items.ListByOrderID(ctx, response.ID)
	if err != nil {
		return nil, errors.Wrap(err, "orders: OrdersService.Create repository.Items.List error")
	}
	response.Items = items
	return response, nil
}

func (o *OrdersServiceImpl) Get(ctx context.Context, id string) (*entity.Order, error) {
	order, err := o.repository.Orders.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "orders: OrdersService.Get repository.Orders.Get error")
	}

	items, err := o.repository.Items.ListByOrderID(ctx, order.ID)
	if err != nil {
		return nil, errors.Wrap(err, "orders: OrdersService.Get repository.Items.List error")
	}
	order.Items = items
	return order, nil
}

func (o *OrdersServiceImpl) List(ctx context.Context) ([]*entity.Order, error) {
	orders, err := o.repository.Orders.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "orders: OrdersService.List repository.Orders.List error")
	}

	return orders, nil
}

func (o *OrdersServiceImpl) TextToSpeach(ctx context.Context, text string) ([]byte, error) {
	client := openai.NewClient(option.WithAPIKey("sk-6G8jjZyJXKDLbGelM1CCm_tPmIzrBMdjLCAh8O3RnbT3BlbkFJG5U2e61mJVJOWgV5tuRGid6WTFvdfDVlI0Tz1jbQcA"))
	response, err := client.Audio.Speech.New(ctx, openai.AudioSpeechNewParams{
		Model:          openai.F(openai.SpeechModelTTS1),
		Input:          openai.String(text),
		ResponseFormat: openai.F(openai.AudioSpeechNewParamsResponseFormatMP3),
		Voice:          openai.F(openai.AudioSpeechNewParamsVoiceAlloy),
	})

	if err != nil {
		return nil, errors.Wrap(err, "orders: OrdersService.textToSpeach client.Audio.Speech.New error")
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "orders: OrdersService.textToSpeach io.ReadAll error")
	}

	return b, nil
}
