package talkie

import (
	"fmt"
	"sync"
)

type Role struct {
	talk  *Talkie
	Name  string
	color ColorWriter
	in    chan Message
	resp  Responder
}

func (t *Talkie) NewRole(name, color string, resp Responder) *Role {
	role := &Role{
		talk:  t,
		Name:  name,
		color: *NewColorWriter(color),
		in:    make(chan Message),
		resp:  resp,
	}

	if role.Name == username {
		return role
	}
	if !t.add(role) {
		fmt.Printf("Role $s skipped duplicate\n", role.Name)
	}
	return role
}

func (r *Role) Print(msg string) (err error) {
	_, err = r.color.Printf("%s: %s\n", r.Name, msg)
	return err
}

func (r *Role) Send(br *Broker, msg string) {
	br.Send(Message{
		Sender:  r,
		Content: msg,
	})
}

func (r *Role) Start(br *Broker, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for msg := range r.in {
			if r.resp.NeedReponse(msg) {
				response, leave := r.resp.Respond(msg)
				r.talk.Broadcast(Message{
					Content: response,
					Sender:  r,
				})
				if leave {
					r.talk.remove(r)
				}
			}
		}
		wg.Done()
	}()
}

func (r *Role) Recieve(msg Message) {
	if r == msg.Sender {
		return // ignore messages by me
	}
	r.in <- msg
}

func (r *Role) Close() {
	close(r.in)
}
