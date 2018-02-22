package main

import (
	"fmt"
	"time"

	hey "github.com/hatajoe/hey-client"
)

type Receiver struct{}

func (r Receiver) Receive(msg string) {
	fmt.Println(msg)
}

var (
	url    = "ws://localhost:9218/"
	origin = "http://localhost:9218/"
)

func main() {
	r := &Receiver{}
	hey.Connect(url, origin, r)
	hey.Send(1, "Hello, World")
	time.Sleep(time.Second)
	hey.Send(1, "See you!")
	time.Sleep(time.Second * 10)
	hey.Disconnect()
}
