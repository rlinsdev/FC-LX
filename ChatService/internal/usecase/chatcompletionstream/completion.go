package chatcompletionstream

import (
	"context"
	"errors"

	"github.com/rlinsdev/FC-LX/tree/main/ChatService/internal/domain/entity"
	"github.com/rlinsdev/FC-LX/tree/main/ChatService/internal/domain/gateway"
	"github.com/sashabaranov/go-openai"
)

type ChatCompletionUseCase struct {
	chatGateway gateway.ChatGateway
	OpenAiClient *openai.Client
}

type ChatCompletionConfigInputDTO struct {
	Model 				string
	ModelMaxTokenx		int
	Temperature			float32
	TopP				float32
	N					float32
	Stop				[]string
	MaxTokens			int
	PresencePenalty 	float32
	FrequencyPenalty 	float32
	InitialSystemMessage string
}

type ChatCompletionInputDTO struct {
	ChatID			string
	UserID			string
	UserMessage		string
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
	userMessage, err := entity.NewMessage("user", input.UserMessage, chat.Config.Model)
	if err != nil {
		return nil, errors.New("error creating user message:" + err.Error())
	}
	err = chat.AddMessage(userMessage)
	if err != nil {
		return nil, errors.New("error adding new message: " + err.Error())
	}

	messages := []openai.ChatCompletionMessage{}
	for _, msg := range chat.Messages {
		messages = append(messages, openai.ChatCompletionMessage{
			Role: msg.Role,
			Content: msg.Content,
		})
	}
}

func createNewChat(input ChatCompletionInputDTO) (*entity.Chat, error) {
	model := entity.NewModel(input.Config.Model, input.Config.ModelMaxTokenx)
	chatConfig := &entity.Chatconfig{
		Temperature: 		input.Config.Temperature,
		TopP: 				input.Config.TopP,
		N:					int(input.Config.N),
		Stop: 				input.Config.Stop,
		MaxTokens: 			input.Config.MaxTokens,
		PresencePenalty: 	input.Config.PresencePenalty,
		FrequencyPenalty: 	input.Config.FrequencyPenalty,
		Model:				model,
	}
	initialMessage, err := entity.NewMessage("system", input.Config.InitialSystemMessage, model)
	if err != nil {
		return nil, errors.New("Error on creating initial message:" + err.Error())
	}
	chat, err := entity.NewChat(input.UserID, initialMessage, chatConfig)
	if err != nil {
		return nil, errors.New("error creating new chat: "+ err.Error())
	}
	return chat, nil
}