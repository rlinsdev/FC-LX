package entity

import "errors"

type Chatconfig struct {
	Model            *Model
	Temperature      float32  // Precision of answered
	TopP             float32  // Conservative in words used
	N                int      // 1 by default... Messages returned
	Stop             []string // key to stop chat (reserved word)
	MaxTokens        int      // number of token generated
	PresencePenalty  float32  //
	FrequencyPenalty float32  //
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *Chatconfig
}

func (c *Chat) Validate() error {
	if c.UserID == "" {
			return errors.New(("user id is empty"))
	}
}

func (c *Chat) AddMessage(m *Message) error {
	if c.Status == "ended" {
		return errors.New("chat is ended. No more message allowed")
	}
	for {
		// The new message allowed in the free space
		if c.Config.Model.GetMaxToken() >= m.GetQtdTokens()+c.TokenUsage {
			c.Messages = append(c.Messages, m)
			c.RefreshTokenUsage()
			break
		}
		c.ErasedMessages = append(c.ErasedMessages, c.Messages[0])
		c.Messages = c.Messages[1:] //Remove the oldest message
		c.RefreshTokenUsage()
	}
	return nil
}

func (c *Chat) GetMessages() []*Message {
	return c.Messages
}

func (c *Chat) CountMessages() int {
	return len(c.Messages)
}

func (c *Chat) End() {
	c.Status = "ended"
}

func (c *Chat) RefreshTokenUsage() {
	c.TokenUsage = 0
	for m:= range c.Messages {
		c.TokenUsage += c.Messages[m].GetQtdTokens()
	}
}
