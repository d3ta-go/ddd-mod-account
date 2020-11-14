package repository

import (
	"fmt"
	"testing"

	"github.com/d3ta-go/ddd-mod-account/modules/account/domain/repository"
	domSchemaAuth "github.com/d3ta-go/ddd-mod-account/modules/account/domain/schema/auth"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/d3ta-go/system/system/utils"
	"github.com/spf13/viper"
)

func newConfig(t *testing.T) (*config.Config, *viper.Viper, error) {

	c, v, err := config.NewConfig("../../../../conf")
	if err != nil {
		return nil, nil, err
	}
	if !c.CanRunTest() {
		panic(fmt.Sprintf("Cannot Run Test on env `%s`, allowed: %v", c.Environment.Stage, c.Environment.RunTestEnvironment))
	}
	return c, v, nil
}

func newRepo(t *testing.T) (repository.IAuthenticationRepo, *handler.Handler, error) {
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

	r, err := NewAuthenticationRepo(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.DefaultIdentity, identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}
	i.Claims.Username = "test.d3tago"
	i.RequestInfo.Host = "127.0.0.1:2020"

	return i
}

func TestAuthRepo_Register(t *testing.T) {
	repo, h, err := newRepo(t)
	if err != nil {
		t.Errorf("Error.newRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.infra-layer.repository.register")

	unique := utils.GenerateUUID()
	req := &domSchemaAuth.RegisterRequest{
		Username:  fmt.Sprintf(testData["username"], unique),
		Password:  testData["password"],
		Email:     fmt.Sprintf(testData["email"], unique),
		NickName:  testData["nick-name"],
		Captcha:   testData["captcha-value"], // validation on interface
		CaptchaID: testData["captcha-id"],    // validation on interface
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Rew.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Register(req, i)
	if err != nil {
		t.Errorf("Error.AuthRepo.Register: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		// save to test-data
		// save result for next test
		viper.Set("test-data.account.auth.infra-layer.repository.activate-registration.activation-code", res.ActivationCode)
		viper.Set("test-data.account.auth.infra-layer.repository.activate-registration.email", res.Email)

		viper.Set("test-data.account.auth.infra-layer.repository.login.username", req.Username)
		viper.Set("test-data.account.auth.infra-layer.repository.login.password", req.Password)
		viper.Set("test-data.account.auth.infra-layer.repository.login.captcha-value", req.Captcha)
		viper.Set("test-data.account.auth.infra-layer.repository.login.captcha-id", req.CaptchaID)
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("Resp.AuthRepo.Register: %s", string(respJSON))
	}
}

func TestAuthRepo_ActivateRegistration(t *testing.T) {
	repo, h, err := newRepo(t)
	if err != nil {
		t.Errorf("Error.newRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.infra-layer.repository.activate-registration")

	req := &domSchemaAuth.ActivateRegistrationRequest{
		ActivationCode: testData["activation-code"],
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Rew.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.ActivateRegistration(req, i)
	if err != nil {
		t.Errorf("Error.AuthRepo.ActivateRegistration: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.AuthRepo.ActivateRegistration: %s", string(respJSON))
	}
}

func TestAuthRepo_Login(t *testing.T) {
	repo, h, err := newRepo(t)
	if err != nil {
		t.Errorf("Error.newRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.infra-layer.repository.login")

	req := &domSchemaAuth.LoginRequest{
		Username:  testData["username"],
		Password:  testData["password"],
		Captcha:   testData["captcha-value"], // validation on interface
		CaptchaID: testData["captcha-id"],    // validation on interface
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Rew.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Login(req, i)
	if err != nil {
		t.Errorf("Error.AuthRepo.Login: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		// save to test-data
		// save result for next test
		viper.Set("test-data.account.auth.infra-layer.repository.response.session.login.token", res.Token)
		viper.Set("test-data.account.auth.infra-layer.repository.response.session.login.expiret-at", res.ExpiredAt)
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("Resp.AuthRepo.Login: %s", string(respJSON))
	}
}

func TestAuthRepo_LoginApp(t *testing.T) {
	repo, h, err := newRepo(t)
	if err != nil {
		t.Errorf("Error.newRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.infra-layer.repository.login-app")

	req := &domSchemaAuth.LoginAppRequest{
		ClientKey: testData["client-key"],
		SecretKey: testData["secret-key"],
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.LoginApp(req, i)
	if err != nil {
		t.Errorf("Error.AuthRepo.LoginApp: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		// save to test-data
		// save result for next test
		viper.Set("test-data.account.auth.infra-layer.repository.response.session.login-app.token", res.Token)
		viper.Set("test-data.account.auth.infra-layer.repository.response.session.login-app.expiret-at", res.ExpiredAt)
		viper.Set("test-data.account.auth.infra-layer.repository.response.session.login-app.client-app-code", res.ClientAppCode)

		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("Resp.AuthRepo.LoginApp: %s", string(respJSON))
	}
}
