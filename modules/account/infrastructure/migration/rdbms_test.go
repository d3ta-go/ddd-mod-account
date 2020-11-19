package migration

import (
	"testing"

	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/spf13/viper"
)

func newRDBMSMigration(t *testing.T) (*RDBMSMigration, *handler.Handler, error) {
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

	m, err := NewRDBMSMigration(h)
	if err != nil {
		return nil, nil, err
	}

	return m, h, nil
}

func TestRDBMSMigration_Run(t *testing.T) {
	mig, _, err := newRDBMSMigration(t)
	if err != nil {
		t.Errorf("Error.newRDBMSMigration: %s", err.Error())
	}

	if mig != nil {
		tests := []struct {
			name    string
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name:    "Run OK",
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := mig.Run(); (err != nil) != tt.wantErr {
					t.Errorf("RDBMSMigration.Run() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestRDBMSMigration_RollBack(t *testing.T) {
	mig, _, err := newRDBMSMigration(t)
	if err != nil {
		t.Errorf("Error.newRDBMSMigration: %s", err.Error())
	}

	if mig != nil {
		tests := []struct {
			name    string
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name:    "RollBack OK",
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := mig.RollBack(); (err != nil) != tt.wantErr {
					t.Errorf("RDBMSMigration.RollBack() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}
