package chat

import "sync"

type Chatroom struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Members  []int     `json:"members"`
	Messages []Message `json:"messages"`
	mu       sync.Mutex
}

type DM struct {
	ID int `json:"id"`
	
}

func NewChatroom(id int, name string, members []int) *Chatroom {
	return &Chatroom{
		ID:      id,
		Name:    name,
		Members: members,
		mu:      sync.Mutex{},
	}
}

func (c *Chatroom) SendMessage() {
}

func (c *Chatroom) BroadcastMessage() {
}


