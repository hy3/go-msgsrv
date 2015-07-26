# go-msgsrv

## Description
go-msgsrv is a simple on-memory message server.

Use for application test or demonstration.

## Installation
This tool can be installed with the `go get` command.

    go get github.com/hy3/go-msgsrv

## Usage
Start server with listen-port 80(default).

    go-msgsrv

Start server with listen-port 1234.

    go-msgsrv -p 1234

## APIs
### URLs
Send message to drawer named JohnDoe.

    POST /JohnDoe/messages/new

Send broadcast message(Broadcast message is delivered to all drawers.).

    POST /messages/new

Show all messages in drawer named JohnDoe. Message is shown with JSON Format

    GET /JohnDoe/messages



### Request parameters
When request post message, set parameter *from* and *body*

* from
  - Set message sender name.

* body
  - Set message body.

## Author
Takahiro Honda (a.k.a hy3)
