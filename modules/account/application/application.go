package application

import (
	"github.com/d3ta-go/ddd-mod-account/modules/account/application/service"
	"github.com/d3ta-go/system/system/handler"
)

// NewAccountApp new AccountApp
func NewAccountApp(h *handler.Handler) (*AccountApp, error) {
	var err error

	app := new(AccountApp)
	app.handler = h

	if app.AuthenticationSvc, err = service.NewAuthenticationSvc(h); err != nil {
		return nil, err
	}

	return app, nil
}

// AccountApp type
type AccountApp struct {
	handler           *handler.Handler
	AuthenticationSvc *service.AuthenticationSvc
}
