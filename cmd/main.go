package main

import (
	"fmt"

	"github.com/robbyriverside/talkie"
)

func main() {
	t := talkie.New()

	bill := t.NewRole("Bill", talkie.Cyan, &talkie.EchoResponder{})
	t.NewRole("Terry", talkie.Cyan, &talkie.EchoResponder{})
	t.NewRole("Joe", talkie.Cyan, &talkie.EchoResponder{})

	t.Start()

	t.Broadcast(talkie.Message{
		Content: "echo echo",
		Sender:  bill,
	})

	fmt.Println("start waiting")
	t.Wait()
	fmt.Println("done waiting")
}
