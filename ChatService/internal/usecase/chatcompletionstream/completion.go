package chatcompletionstream

import (
	"context"
	"errors"

	"github.com/rlinsdev/FC-LX/tree/main/ChatService/internal/domain/gateway"
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

type ChatCompletionInputDTO struct {
	ChatID			string
	UserID			string
	UserMessage	string
	Config			ChatCompletionConfigInputDTO
}

type ChatCompletionOutputDTO struct {
	ChatID		string
	UserID		string
	Content		string
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openAiClient *openai.Client) *ChatCompletionUseCase {
	return &ChatCompletionUseCase{
		chatGateway: chatGateway,
		OpenAiClient: openAiClient,
	}
}

func (uc *ChatCompletionUseCase) Execute(ctx context.Context, input ChatCompletionInputDTO) (*ChatCompletionOutputDTO, error) {
	// Existent chat or new chat?
	chat, err :=uc.chatGateway.FindChatById(ctx, input.ChatID)
	
	if err != nil {
		if err.Error() == "chat not found" {
			// Create a new chat
			chat, err = createNewChat(input)
			if err != nil {
				return nil, errors.New("Error creating chat:" + err.Error())
			}
			// Save on DB
			err = uc.chatGateway.CreateChat(ctx, chat)
			if err != nil {
				return nil, errors.New("error on saving new Chat: "+err.Error())
			}
		} else {
			return nil, errors.New("error fetchin existing chat: "+ err.Error())
		}
	}
}