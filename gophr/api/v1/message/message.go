package message

import (
	"bytes"
	"text/template"
)

var msgEn = `
{{ define "UsernameExists"}} Username already exists {{ end }}
{{ define "EmailExists"}} Email already exists {{ end }}
`

func New(language string) *Message {
	t := template.New(language)
	t, _ = t.Parse(msgEn)
	return &Message{
		t: t,
		params: make(map[string]interface{}),
	}
}

type Message struct {
	t      *template.Template
	name   string
	params map[string]interface{}
}

func (m *Message) AddName(name string) *Message {
	m.name = name
	return m
}

func (m *Message) AddParams(id string, params ...interface{}) *Message {
	for _, p := range params {
		m.params[id] = append(params, p)
	}
	return m
}

func (m *Message) Apply() string {
	var buff bytes.Buffer
	_ = m.t.ExecuteTemplate(&buff, m.name, m.params)
	return buff.String()
}

