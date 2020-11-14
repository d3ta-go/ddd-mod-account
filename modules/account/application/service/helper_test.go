package service

import (
	"fmt"
	"testing"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/context"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
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
	c.IAM.Casbin.ModelPath = "../../../../conf/casbin/casbin_rbac_rest_model.conf"
	return c, v, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	j, err := identity.NewJWT(h)
	if err != nil {
		return identity.Identity{}
	}
	claims, token, _, err := j.GenerateAnonymousToken()
	if err != nil {
		return identity.Identity{}
	}
	i, err := identity.NewIdentity(identity.DefaultIdentity, identity.TokenJWT, token, claims, context.NewCtx(context.SystemCtx), h)
	if err != nil {
		return identity.Identity{}
	}
	return i
}
