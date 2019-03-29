package Validation

import (
	"encoding/json"
	"github.com/apmath-web/expenses/Domain"
)

type Validation struct {
	message  string
	messages []Domain.MessageInterface
}

func (v *Validation) AddMessage(message Domain.MessageInterface) {
	v.messages = append(v.messages, message)
}

func GenMessage(field, text string) *Message {
	m := new(Message)
	m.field = field
	m.text = text
	return m
}

type Message struct {
	text, field string
}

func (m *Message) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		m.field: m.text,
	})
}
