package chatcompletionstream

import (
	"github.com/rlinsdev/fclx/chatservice/internal/domain/gateway"
	"github.com/sashabaranov/go-openai"
)

type ChatCompletionUseCase struct {
	chatGateway gateway.ChatGateway
	OpenAiClient *openai.Client
}

type ChatCompletionConfigInputDTO struct {
	Mode 							string
	ModelMaxTokenx		int
	Temperature				float32
	TopP							float32
	N									float32
	Stop							[]string
	MaxTokens					int
	PresencePenalty 	float32
	FrequencyPenalty 	float32
	InitialSystemMessage string
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openAiClient *openai.Client) *ChatCompletionUseCase {
	return &ChatCompletionUseCase{
		chatGateway: chatGateway,
		OpenAiClient: openAiClient,
	}
}