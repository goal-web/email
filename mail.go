package email

import "github.com/goal-web/contracts"

func New(subject string, content contracts.EmailContent) contracts.Mailable {
	return &Mail{
		Subject: subject,
		Text:    content.Text(),
		Html:    content.Html(),
		To:      make([]string, 0),
		Cc:      make([]string, 0),
		Bcc:     make([]string, 0),
	}
}

type Mail struct {
	From    string
	Subject string
	Text    string
	Html    string
	To      []string
	Cc      []string
	Bcc     []string
	queue   string
	delay   int
}

func (mail *Mail) SetCc(address ...string) contracts.Mailable {
	mail.Cc = address
	return mail
}

func (mail *Mail) SetBcc(address ...string) contracts.Mailable {
	mail.Bcc = address
	return mail
}

func (mail *Mail) SetTo(address ...string) contracts.Mailable {
	mail.To = address
	return mail
}

func (mail *Mail) Queue(queue string) contracts.Mailable {
	mail.queue = queue
	return mail
}

func (mail *Mail) Delay(delay int) contracts.Mailable {
	mail.delay = delay
	return mail
}

func (mail *Mail) GetCc() []string {
	return mail.Cc
}

func (mail *Mail) GetBcc() []string {
	return mail.Bcc
}

func (mail *Mail) GetTo() []string {
	return mail.To
}

func (mail *Mail) GetSubject() string {
	return mail.Subject
}

func (mail *Mail) SetFrom(from string) contracts.Mailable {
	mail.From = from
	return mail
}
func (mail *Mail) GetFrom() string {
	return mail.From
}

func (mail *Mail) GetText() string {
	return mail.Text
}

func (mail *Mail) GetHtml() string {
	return mail.Html
}

func (mail *Mail) GetQueue() string {
	return mail.queue
}

func (mail *Mail) GetDelay() int {
	return mail.delay
}
