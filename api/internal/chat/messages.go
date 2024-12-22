package chat

import (
	"time"
)

type Message struct {
	User      int       `json:"user"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}

func NewMessage(user int, body string) *Message {
	return &Message{
		User:      user,
		Body:      body,
		Timestamp: time.Now(),
	}
}
