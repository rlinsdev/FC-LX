package entity

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
