package hey

import (
	"log"

	broadcaster "github.com/dev-cloverlab/go-message-broadcaster"

	"golang.org/x/net/websocket"
)

var (
	ws *websocket.Conn
)

type OnReceive interface {
	Receive(string)
}

func Connect(url, origin string, h OnReceive) {
	var err error
	ws, err = websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatalf("websocket.Dial failed: %s", err)
	}

	go receive(h)
}

func Disconnect() {
	if err := ws.Close(); err != nil {
		log.Fatalf("ws.Close failed: %s", err)
	}
}

func Send(hid int, buf string) {
	msg := &broadcaster.RequestMessage{
		HandlerID: broadcaster.MessageHandlerID(hid),
		Body:      []byte(buf),
	}
	if err := websocket.JSON.Send(ws, msg); err != nil {
		log.Printf("websocket.JSON.Send failed: %s", err)
	}
}

func receive(h OnReceive) {
	var msg broadcaster.ResponseMessage
	for {
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			log.Printf("websocket.JSON.Receive failed: %s", err)
			continue
		}
		h.Receive(string(msg.Body))
	}
}
