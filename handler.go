package main

import (
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

	return router
}

func addBroadcastMessage(writer http.ResponseWriter, request *http.Request) {
	msg := message.New(request.FormValue("from"), message.Broadcast, request.FormValue("body"))
	msgBox.Post(msg)
}

func addMessage(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	name := params["name"]
	msg := message.New(request.FormValue("from"), name, request.FormValue("body"))
	msgBox.Post(msg)
}

func showMessages(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	name := params["name"]
	messages := msgBox.Pickup(name)
	if messages == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	json, err := message.ConvertToJSON(messages)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(json)
}
