package message

import (
	"testing"
)

func TestPickup(t *testing.T) {
	msgBox := NewMessageBox()
	msgBox.Drawers["JohnDoe"] = NewDrawer()
	msgBox.Drawers["JohnDoe"].appendMessage(New("someone", "JohnDoe", "testmsg1"))
	msgBox.Drawers["JohnDoe"].appendMessage(New("someone", "JohnDoe", "testmsg2"))

	messages := msgBox.Pickup("JohnDoe")
	if len(messages) != 2 {
		t.Fatalf("len(messages) => %d, wants %d", len(messages), 2)
	}
	if messages[0].Body != "testmsg1" {
		t.Errorf("messages[0].Body => %s, wants %s", messages[0].Body, "testmsg1")
	}
	if messages[1].Body != "testmsg2" {
		t.Errorf("messages[1].Body => %s, wants %s", messages[1].Body, "testmsg2")
	}
	if len(msgBox.Drawers["JohnDoe"].Messages) != 0 {
		t.Error("Message must be empty after Pickup, but is not.")
	}
}
