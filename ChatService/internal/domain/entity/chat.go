package entity

type Chatconfig struct {
	Model *Model
}

type Chat struct {
	ID         string
	UserID     string
	Status     string
	TokenUsage int
	Config     *Chatconfig
}
