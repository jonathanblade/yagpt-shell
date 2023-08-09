package domain

import "time"

type Role struct {
	Name string
}

var (
	User      Role = Role{Name: "Пользователь"}
	Assistant      = Role{Name: "Ассистент"}
)

type Message struct {
	Role Role
	Text string
	Time time.Time
}

func NewUserMessage(text string) Message {
	return Message{Role: User, Text: text, Time: time.Now()}
}

func NewAssistantMessage(text string) Message {
	return Message{Role: Assistant, Text: text, Time: time.Now()}
}
