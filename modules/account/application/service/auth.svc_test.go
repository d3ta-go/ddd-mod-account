package service

import (
	"encoding/json"
	"fmt"
	"testing"

	appDTOAuth "github.com/d3ta-go/ddd-mod-account/modules/account/application/dto/auth"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/d3ta-go/system/system/utils"
	"github.com/spf13/viper"
)

func newAuthenticationSvc(t *testing.T) (*AuthenticationSvc, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, v, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetDefaultConfig(c)
	h.SetViper("config", v)

	// viper for test-data
	viperTest := viper.New()
	viperTest.SetConfigType("yaml")
	viperTest.SetConfigName("test-data")
	viperTest.AddConfigPath("../../../../conf/data")
	viperTest.ReadInConfig()
	h.SetViper("test-data", viperTest)

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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.register")

	unique := utils.GenerateUUID()
	req := appDTOAuth.RegisterReqDTO{}
	req.Username = fmt.Sprintf(testData["username"], unique)
	req.Password = testData["password"]
	req.Email = fmt.Sprintf(testData["email"], unique)
	req.NickName = testData["nick-name"]
	req.Captcha = testData["captcha-value"] // validation on interface
	req.CaptchaID = testData["captcha-id"]  // validation on interface

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
		// save to test-data
		// save result for next test
		viper.Set("test-data.account.auth.activate-registration.activation-code", resp.ActivationCode)
		viper.Set("test-data.account.auth.activate-registration.email", resp.Email)

		viper.Set("test-data.account.auth.login.username", req.Username)
		viper.Set("test-data.account.auth.login.password", req.Password)
		viper.Set("test-data.account.auth.login.captcha-value", req.Captcha)
		viper.Set("test-data.account.auth.login.captcha-id", req.CaptchaID)
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.activate-registration")

	req := appDTOAuth.ActivateRegistrationReqDTO{}
	req.ActivationCode = testData["activation-code"]

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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.login")

	req := appDTOAuth.LoginReqDTO{}
	req.Username = testData["username"]
	req.Password = testData["password"]
	req.Captcha = testData["captcha-value"] // validation on interface
	req.CaptchaID = testData["captcha-id"]  // validation on interface

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
		// save to test-data
		// save result for next test
		viper.Set("test-data.account.auth.response.session.login.token", resp.Token)
		viper.Set("test-data.account.auth.response.session.login.expiret-at", resp.ExpiredAt)
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.login-app")

	req := appDTOAuth.LoginAppReqDTO{}
	req.ClientKey = testData["client-key"]
	req.SecretKey = testData["secret-key"]

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
		// save to test-data
		// save result for next test
		viper.Set("test-data.account.auth.response.session.login-app.token", resp.Token)
		viper.Set("test-data.account.auth.response.session.login-app.expiret-at", resp.ExpiredAt)
		viper.Set("test-data.account.auth.response.session.login-app.client-app-code", resp.ClientAppCode)

		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}
