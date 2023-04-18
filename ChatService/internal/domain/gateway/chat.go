package gateway

import (
	"context"

	"github.com/rlinsdev/fclx/chatservice/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	FindChatById(ctx context.Context, chatID string) (*entity.Chat, error)
	SaveChage(ctx context.Context, chat *entity.Chat) error
}