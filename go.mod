module github.com/goal-web/email

go 1.17

require (
	github.com/goal-web/contracts v0.1.34
	github.com/goal-web/supports v0.1.14
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
)

require (
	github.com/apex/log v1.9.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
)

replace github.com/goal-web/contracts => ../contracts
