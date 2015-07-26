package message

import (
	"time"
)

// MessageBox is a message box which holds all messages in its drawers.
type MessageBox struct {
	Drawers map[string]*Drawer
}

// NewMessageBox creates a new MessageBox object with no Drawers.
func NewMessageBox() *MessageBox {
	m := new(MessageBox)
	m.Drawers = make(map[string]*Drawer)
	return m
}

// Drawer is a drawer of message box.
// Every drawer holds messages for someone.
type Drawer struct {
	Messages []*Message `json:"messages"`
}

// NewDrawer creates a new Drawer object with no Messages.
func NewDrawer() *Drawer {
	d := new(Drawer)
	d.Messages = make([]*Message, 0, 10)
	return d
}

// Message is a message sent from someone to someone.
// If value of To is "all", it is a broadcast message.
type Message struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}

// New creates a new Message object.
func New(from, to, body string) *Message {
	m := new(Message)
	m.From = from
	m.To = to
	m.Body = body
	m.Timestamp = time.Now()
	return m
}
