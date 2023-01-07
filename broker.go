package talkie

import (
	"fmt"
)

type Broker struct {
	ch chan Message
}

func NewBroker() *Broker {
	return &Broker{
		ch: make(chan Message),
	}
}

func (br *Broker) Send(msg Message) {
	br.ch <- msg
}

func (br *Broker) Start() {
	fmt.Println("Starting conversation:")

	go func() {
		for msg := range br.ch {
			msg.Print()
		}
	}()
}

func (br *Broker) Close() {
	close(br.ch)
}
