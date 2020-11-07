package application

import (
	"reflect"
	"testing"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newHandler(t *testing.T) (*handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetDefaultConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, err
	}

	return h, nil
}

func TestNewAccountApp(t *testing.T) {
	h, err := newHandler(t)
	if err != nil {
		t.Errorf("newHandler: %s", err.Error())
		return
	}

	if h != nil {

		type args struct {
			h *handler.Handler
		}
		tests := []struct {
			name    string
			args    args
			want    *AccountApp
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name:    "Create NewAccountApp",
				args:    args{h: h},
				want:    &AccountApp{handler: h},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := NewAccountApp(tt.args.h)
				if (err != nil) != tt.wantErr {
					t.Errorf("NewAccountApp() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got.handler, tt.want.handler) {
					t.Errorf("NewAccountApp() = %v, want %v", got.handler, tt.want.handler)
				}
			})
		}
	}
}
