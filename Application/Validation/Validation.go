package Validation

import "github.com/apmath-web/expenses/Domain"

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
