package email

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/utils"
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
		config := factory.config.Mailers[name]
		if config == nil {
			panic(Exception{Exception: exceptions.New("factory.getMailer: mailer does not exist", config)})
		}

		if driver, ok := factory.drivers[utils.GetStringField(config, "driver")]; ok {
			factory.mailers[name] = driver(name, config)
		} else {
			panic(Exception{Exception: exceptions.New("factory.getMailer: driver does not exist", config)})
		}
	}

	return factory.mailers[name]
}

func (factory *Factory) Extend(name string, driver contracts.MailerDriver) {
	factory.drivers[name] = driver
}
