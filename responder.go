package talkie

// Responder is a model for generating text responses to input text.
type Responder interface {
	// NeedResponse determines if the input needs a response.
	NeedReponse(in Message) bool
	// Respond take the input and generates an appropriate response.
	// If bool is true then leave the conversation
	Respond(in Message) (string, bool)
}

type EchoResponder struct {
	last string
}

func (er *EchoResponder) NeedReponse(in Message) bool {
	if in.Content != er.last {
		er.last = in.Content
		return true
	}
	return false
}

func (er *EchoResponder) Respond(in Message) (string, bool) {
	return in.Content, true
}
