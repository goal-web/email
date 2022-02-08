package email

import "github.com/goal-web/contracts"

type Config struct {
	Default string
	Mailers map[string]contracts.Fields
}
