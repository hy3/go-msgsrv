package message

import (
	"time"
)

// MessageBox is a message box which holds all messages in its drawers.
type MessageBox struct {
	Drawers map[string]*Drawer
}

// Drawer is a drawer of message box.
// Every drawer holds messages for someone.
type Drawer struct {
	Messages []*Message `json:"messages"`
}

// Message is a message sent from someone to someone.
// If value of To is "all", it is a broadcast message.
type Message struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}
