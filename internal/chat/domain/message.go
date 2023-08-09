package domain

import "time"

type RoleType = string

const (
	UserRole      RoleType = "Пользователь"
	AssistantRole RoleType = "Ассистент"
)

type Message struct {
	Role RoleType
	Text string
	Time time.Time
}

func NewUserMessage(text string) Message {
	return Message{Role: UserRole, Text: text, Time: time.Now()}
}

func NewAssistantMessage(text string) Message {
	return Message{Role: AssistantRole, Text: text, Time: time.Now()}
}
