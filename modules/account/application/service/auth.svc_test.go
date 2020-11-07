package service

import (
	"encoding/json"
	"testing"

	appDTOAuth "github.com/d3ta-go/ddd-mod-account/modules/account/application/dto/auth"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
)

func newAuthenticationSvc(t *testing.T) (*AuthenticationSvc, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetDefaultConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, nil, err
	}

	r, err := NewAuthenticationSvc(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func TestAuthenticationService_Register(t *testing.T) {
	svc, h, err := newAuthenticationSvc(t)
	if err != nil {
		t.Errorf("newAuthenticationSvc: %s", err.Error())
		return
	}

	req := appDTOAuth.RegisterReqDTO{}
	req.Username = "admin.d3tago"
	req.Password = "P4s$W0rd!@!"
	req.Email = "admin.d3tago@email.com"
	req.NickName = "Hari"
	req.Captcha = "just-capthcha-value" // validation on interface
	req.CaptchaID = "just-chaptcha-id"  // validation on interface

	i := newIdentity(h, t)

	resp, err := svc.Register(&req, i)
	if err != nil {
		t.Errorf("Register: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}

func TestAuthenticationService_ActivateRegistration(t *testing.T) {
	svc, h, err := newAuthenticationSvc(t)
	if err != nil {
		t.Errorf("newAuthenticationSvc: %s", err.Error())
		return
	}

	req := appDTOAuth.ActivateRegistrationReqDTO{}
	req.ActivationCode = "a70112cc-bca6-45c2-9bb6-cf3a56daf566"

	i := newIdentity(h, t)

	resp, err := svc.ActivateRegistration(&req, i)
	if err != nil {
		t.Errorf("ActivateRegistration: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}

func TestAuthenticationService_Login(t *testing.T) {
	svc, h, err := newAuthenticationSvc(t)
	if err != nil {
		t.Errorf("newAuthenticationSvc: %s", err.Error())
		return
	}

	req := appDTOAuth.LoginReqDTO{}
	req.Username = "admin.d3tago"
	req.Password = "P4s$W0rd!@!"
	req.Captcha = "just-capthcha-value" // validation on interface
	req.CaptchaID = "just-chaptcha-id"  // validation on interface

	i := newIdentity(h, t)

	resp, err := svc.Login(&req, i)
	if err != nil {
		t.Errorf("Login: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}

func TestAuthenticationService_LoginApp(t *testing.T) {
	svc, h, err := newAuthenticationSvc(t)
	if err != nil {
		t.Errorf("newAuthenticationSvc: %s", err.Error())
		return
	}

	req := appDTOAuth.LoginAppReqDTO{}
	req.ClientKey = "53102ba5-b6b2-47ad-a68d-682463a8be29"
	req.SecretKey = "OTk5ZDlmYjJlZGUyMjAxNTZkZThiNmNkMmJmNDI1NjdiNTYzMzcxNDEwNzNiNDBjM2NhZmIxOWY3NzZmYzhmNg=="

	i := newIdentity(h, t)

	resp, err := svc.LoginApp(&req, i)
	if err != nil {
		t.Errorf("LoginApp: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}
