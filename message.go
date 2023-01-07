package talkie

type Message struct {
	Content string
	Sender  *Role
}

func (msg Message) Print() error {
	return msg.Sender.Print(msg.Content)
}
