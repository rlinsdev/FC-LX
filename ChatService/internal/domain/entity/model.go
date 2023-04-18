package entity

type Model struct {
	Name string
	MaxTokens int
}

func NewModel(name string, MaxTokens int) *Model {
	return &Model {
		Name: name,
		MaxTokens: MaxTokens,
	}
}

func (m *Model) GetMaxToken() int {
	return m.MaxTokens
}

func (m *Model) GetModelName() string {
	return m.Name
}