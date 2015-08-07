package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hy3/go-msgsrv/message"
)

var msgBox = message.NewMessageBox()

const (
	methodGet    = "GET"
	methodPost   = "POST"
	methodPut    = "PUT"
	methodDelete = "DELETE"
)

func SetupHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/messages/new", addBroadcastMessage).Methods(methodPost)
	router.HandleFunc("/{name:[0-9a-zA-Z]+}/messages/new", addMessage).Methods(methodPost)
	router.HandleFunc("/{name:[0-9a-zA-Z]+}/messages", showMessages).Methods(methodGet)
	router.HandleFunc("/dump", dump).Methods(methodGet)

	return router
}

func addBroadcastMessage(writer http.ResponseWriter, request *http.Request) {
	msg := message.New(request.FormValue("from"), message.Broadcast, request.FormValue("body"))
	fmt.Printf("Broadcast message received from %s.\n", msg.From)
	if err := msgBox.Post(msg); err != nil {
		fmt.Printf("Broadcast message post failed: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func addMessage(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	name := params["name"]
	msg := message.New(request.FormValue("from"), name, request.FormValue("body"))
	fmt.Printf("Message received from %s to %s.\n", msg.From, msg.To)
	if err := msgBox.Post(msg); err != nil {
		fmt.Printf("Message post failed: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func showMessages(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	name := params["name"]
	fmt.Printf("Request received: pickup messages in %s.\n", name)
	messages := msgBox.Pickup(name)
	if messages == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	json, err := message.ConvertToJSON(messages)
	if err != nil {
		fmt.Printf("JSON error: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(json)
}

func dump(writer http.ResponseWriter, request *http.Request) {
	json, err := msgBox.Dump()
	if err != nil {
		fmt.Printf("JSON error: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(json)
}
