package email

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"net/smtp"
)

type ServiceProvider struct {
	app contracts.Application
}

func (service *ServiceProvider) Register(application contracts.Application) {
	service.app = application

	application.Singleton("mail.factory", func(config contracts.Config, queue contracts.Queue) contracts.EmailFactory {
		return &Factory{
			config:  config.Get("mail").(Config),
			mailers: map[string]contracts.Mailer{},
			drivers: map[string]contracts.MailerDriver{
				"mailer": func(name string, config contracts.Fields) contracts.Mailer {
					return &Mailer{
						name: name,
						from: utils.GetStringField(config, "from"),
						auth: smtp.PlainAuth(
							utils.GetStringField(config, "identity"),
							utils.GetStringField(config, "username"),
							utils.GetStringField(config, "password"),
							utils.GetStringField(config, "host"),
						),
						address: fmt.Sprintf("%s:%s", utils.GetStringField(config, "host"), utils.GetStringField(config, "port")),
						queue:   queue,
					}
				},
			},
		}
	})

	application.Singleton("mailer", func(factory contracts.EmailFactory) contracts.Mailer {
		return factory.Mailer()
	})
}

func (service *ServiceProvider) Start() error {
	service.app.Call(func(serializer contracts.ClassSerializer) {
		serializer.Register(JobClass)
	})
	return nil
}

func (service *ServiceProvider) Stop() {
}
