package email

import "github.com/goal-web/contracts"

func Text(text string) contracts.EmailContent {
	return Textual{text: text}
}

type Textual struct {
	text string
}

func (t Textual) Text() string {
	return t.text
}

func (t Textual) Html() string {
	return t.text
}
