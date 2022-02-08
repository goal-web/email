package email

import (
	"github.com/goal-web/contracts"
)

type Factory struct {
	config  Config
	mailers map[string]contracts.Mailer
	drivers map[string]contracts.MailerDriver
}

func (factory *Factory) Mailer(name ...string) contracts.Mailer {
	mailer := factory.config.Default
	if len(name) > 0 {
		mailer = name[0]
	}

	return factory.getMailer(mailer)
}

func (factory *Factory) getMailer(name string) contracts.Mailer {
	if factory.mailers[name] == nil {

	}

	return factory.mailers[name]
}

func (factory *Factory) Extend(name string, driver contracts.MailerDriver) {
	//TODO implement me
	panic("implement me")
}
