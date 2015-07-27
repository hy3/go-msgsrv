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

Send broadcast message. Broadcast message is delivered to all drawers.

    POST /messages/new

Show and remove all messages in drawer named JohnDoe, with JSON Format. 

    GET /JohnDoe/messages



### Request parameters
Set parameters named *from* and *body* when client sends POST request.

* from
  - Message sender name.

* body
  - Message body.

### Response format
go-msgsvr outputs JSON message when client sends GET request.
Like this,

    [
      {
        "from":"JaneSmith",
        "to":"JohnDoe",
        "body":"Hello, John!",
        "timestamp":"2015-07-27 08:30:15"
      },
      {
        "from":"Administrator",
        "to":"JohnDoe",
        "body":"Software update completed.",
        "timestamp":"2015-07-27 08:32:46"
      }
    ]

## Author
Takahiro Honda (a.k.a hy3)
