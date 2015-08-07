package message

import (
	"fmt"
	"testing"
)

func TestConvertToJSON(t *testing.T) {
	expectedPattern := `[{"from":"a","to":"b","body":"testmsg1","timestamp":"%s"},{"from":"b","to":"c","body":"testmsg2","timestamp":"%s"}]`

	msg1 := New("a", "b", "testmsg1")
	msg2 := New("b", "c", "testmsg2")
	expected := fmt.Sprintf(expectedPattern, msg1.Timestamp, msg2.Timestamp)

	messages := []*Message{msg1, msg2}
	json, err := ConvertToJSON(messages)
	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if string(json) != expected {
		t.Log("ConvertToJSON result is not expected value.")
		t.Log("Expected:")
		t.Log(expected)
		t.Log("Actual:")
		t.Log(json)
		t.Fail()
	}
}

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

func TestPost(t *testing.T) {
	msgBox := NewMessageBox()

	msg1 := New("a", "b", "testmsg1")
	if err := msgBox.Post(msg1); err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	msg2 := New("b", "c", "testmsg2")
	if err := msgBox.Post(msg2); err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	msg3 := New("c", "b", "testmsg3")
	if err := msgBox.Post(msg3); err != nil {
		t.Fatalf("Error occured: %s", err)
	}

	if _, exists := msgBox.Drawers["a"]; exists {
		t.Fatal(`Drawers["a"] must not exist, but it exists.`)
	}

	b, exists := msgBox.Drawers["b"]
	if !exists {
		t.Fatal(`Drawers["b"] must be exist, but dose not exist.`)
	}
	if len(b.Messages) != 2 {
		t.Fatalf("len(b.Messages) => %d, want %d", len(b.Messages), 2)
	}
	if b.Messages[0].Body != "testmsg1" {
		t.Errorf("b.Messages[0].Body => %s, want %s", b.Messages[0].Body, "testmsg1")
	}
	if b.Messages[1].Body != "testmsg3" {
		t.Errorf("b.Messages[1].Body => %s, want %s", b.Messages[1].Body, "testmsg3")
	}

	c, exists := msgBox.Drawers["c"]
	if !exists {
		t.Fatal(`Drawers["c"] must be exist, but dose not exist.`)
	}
	if len(c.Messages) != 1 {
		t.Fatalf("len(c.Messages) => %d, want %d", len(c.Messages), 1)
	}
	if c.Messages[0].Body != "testmsg2" {
		t.Errorf("c.Messages[0].Body => %s, want %s", c.Messages[0].Body, "testmsg2")
	}
}

func TestPost_Broadcast(t *testing.T) {
	msgBox := NewMessageBox()
	msgBox.addDrawer("a")
	msgBox.addDrawer("b")
	msgBox.addDrawer("c")

	msg := New("b", Broadcast, "testmsg")
	msgBox.Post(msg)

	a := msgBox.Drawers["a"]
	if len(a.Messages) != 1 {
		t.Fatalf("len(a.Messages) => %d, want %d", len(a.Messages), 1)
	}
	if a.Messages[0].Body != "testmsg" {
		t.Errorf("a.Messages[0].Body => %s, want %s", a.Messages[0].Body, "testmsg")
	}

	b := msgBox.Drawers["b"]
	if len(b.Messages) != 0 {
		t.Fatalf("len(b.Messages) => %d, want %d", len(b.Messages), 0)
	}

	c := msgBox.Drawers["c"]
	if len(c.Messages) != 1 {
		t.Fatalf("len(c.Messages) => %d, want %d", len(c.Messages), 1)
	}
	if c.Messages[0].Body != "testmsg" {
		t.Errorf("c.Messages[0].Body => %s, want %s", c.Messages[0].Body, "testmsg")
	}
}
