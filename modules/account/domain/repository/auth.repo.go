package repository

import (
	domSchemaAuth "github.com/d3ta-go/ddd-mod-account/modules/account/domain/schema/auth"
	"github.com/d3ta-go/system/system/identity"
)

// IAuthenticationRepo represent AuthenticationRepo Interface
type IAuthenticationRepo interface {
	Register(req *domSchemaAuth.RegisterRequest, i identity.Identity) (*domSchemaAuth.RegisterResponse, error)
	ActivateRegistration(req *domSchemaAuth.ActivateRegistrationRequest, i identity.Identity) (*domSchemaAuth.ActivateRegistrationResponse, error)
	Login(req *domSchemaAuth.LoginRequest, i identity.Identity) (*domSchemaAuth.LoginResponse, error)
	LoginApp(req *domSchemaAuth.LoginAppRequest, i identity.Identity) (*domSchemaAuth.LoginAppResponse, error)
}
