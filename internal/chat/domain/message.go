package domain

import "time"

const (
	UserRole      = "Пользователь"
	AssistantRole = "Ассистент"
)

type Message struct {
	Role string
	Text string
	Time time.Time
}

func NewUserMessage(text string) Message {
	return Message{Role: UserRole, Text: text, Time: time.Now()}
}

func NewAssistantMessage(text string) Message {
	return Message{Role: AssistantRole, Text: text, Time: time.Now()}
}
