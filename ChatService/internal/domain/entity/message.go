package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        string
	Role      string
	Content   string
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}

func NewMessage(role string, content string, model *Model) (*Message, error) {

	msg := &Message {
		ID: 			uuid.New().String(),
		Role: 		role,
		Content: 	content,
		Model:		model,
	}
	return msg, nil
}

func (m *Message) Validate() error {
	if m.Role != "user" && m.Role != "system" && m.Role != "assistant" {
		return errors.New("Invalid Role")
	}
	if m.Content == "" {
		return errors.New("content is empty")
	}
	if m.CreatedAt.IsZero() {
		return errors.New("Invalid created at")
	}
	
	return nil
}