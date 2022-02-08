package email

import (
	"crypto/tls"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"
)

type Mailer struct {
	name      string
	auth      smtp.Auth
	from      string
	address   string
	queue     contracts.Queue
	tlsConfig *tls.Config
}

func (this *Mailer) Raw(subject, text string, to []string) error {
	newEmail := email.NewEmail()
	newEmail.From = this.from
	newEmail.To = to
	newEmail.Subject = subject
	newEmail.Text = []byte(text)
	newEmail.HTML = []byte(text)
	return newEmail.Send(this.address, this.auth)
}

func (this *Mailer) Send(mail contracts.Mailable) error {
	if mail.GetQueue() != "" {
		return this.Queue(mail, mail.GetQueue())
	}
	newEmail := email.NewEmail()

	newEmail.From = mail.GetFrom()

	if newEmail.From == "" {
		newEmail.From = this.from
	}

	newEmail.To = mail.GetTo()
	newEmail.Cc = mail.GetCc()
	newEmail.Bcc = mail.GetBcc()
	newEmail.Subject = mail.GetSubject()
	newEmail.Text = []byte(mail.GetText())
	newEmail.HTML = []byte(mail.GetHtml())

	if this.tlsConfig != nil {
		return newEmail.SendWithStartTLS(this.address, this.auth, this.tlsConfig)
	}

	return newEmail.Send(this.address, this.auth)
}

func (this *Mailer) Queue(mail contracts.Mailable, queue ...string) error {
	if mail.GetDelay() > 0 {
		return this.Later(mail.GetDelay(), mail, queue...)
	}

	return this.queue.Push(&Job{
		UUID:      fmt.Sprintf("email:%s-%s", utils.RandStr(10), mail.GetSubject()),
		CreatedAt: time.Now().Unix(),
		Mail:      ConvertToMail(mail),
	})
}

func (this *Mailer) Later(delay int, mail contracts.Mailable, queue ...string) error {
	return this.queue.Later(time.Now().Add(time.Duration(delay)*time.Second), &Job{
		UUID:      fmt.Sprintf("email:%s-%s", utils.RandStr(10), mail.GetSubject()),
		CreatedAt: time.Now().Unix(),
		Mail:      ConvertToMail(mail),
	}, queue...)
}
