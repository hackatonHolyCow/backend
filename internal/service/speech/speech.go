package speech

import (
	"context"
	"hackathon/backend/pkg/errors"
	"io"

	"github.com/openai/openai-go"
)

type SpeechService interface {
	TextToSpeach(ctx context.Context, text string) ([]byte, error)
}

type SpeechServiceImpl struct {
	client *openai.Client
}

func New(client *openai.Client) SpeechService {
	return &SpeechServiceImpl{
		client: client,
	}
}

func (s *SpeechServiceImpl) TextToSpeach(ctx context.Context, text string) ([]byte, error) {
	response, err := s.client.Audio.Speech.New(ctx, openai.AudioSpeechNewParams{
		Model:          openai.F(openai.SpeechModelTTS1),
		Input:          openai.String(text),
		ResponseFormat: openai.F(openai.AudioSpeechNewParamsResponseFormatMP3),
		Voice:          openai.F(openai.AudioSpeechNewParamsVoiceOnyx),
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
