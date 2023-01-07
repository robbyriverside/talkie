package talkie

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

const username = "User"

var exitMap = map[string]bool{
	"BYE":  true,
	"QUIT": true,
	"exit": true,
}

type Talkie struct {
	broker *Broker
	roles  map[string]*Role
	group  sync.WaitGroup
	stdin  *bufio.Reader
}

func New() *Talkie {
	return &Talkie{
		broker: NewBroker(),
		roles:  make(map[string]*Role),
		stdin:  bufio.NewReader(os.Stdin),
	}
}

func (t *Talkie) Start() {
	t.broker.Start()
	for _, role := range t.roles {
		role.Start(t.broker, &t.group)
	}
	t.ReadInput()
}

func (t *Talkie) ReadInput() {
	user := t.NewRole(username, White, nil)
	for {
		text, _ := t.stdin.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if exitMap[text] {
			t.Close()
			return
		}
		t.Broadcast(Message{
			Sender:  user,
			Content: text,
		})
	}
}

func (t *Talkie) Broadcast(msg Message) {
	t.broker.Send(msg)
	for _, role := range t.roles {
		role.Recieve(msg)
	}
}

func (t *Talkie) Close() {
	for _, role := range t.roles {
		role.Close()
	}
	t.broker.Close()
}

func (t *Talkie) Wait() {
	t.group.Wait()
}

func (t *Talkie) add(role *Role) bool {
	_, found := t.roles[role.Name]
	if found {
		return false
	}
	t.roles[role.Name] = role
	return true
}

func (t *Talkie) remove(role *Role) {
	delete(t.roles, role.Name)
	role.Close()
}
