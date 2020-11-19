package rdbms

import (
	"fmt"

	"github.com/d3ta-go/system/system/handler"
	migRDBMS "github.com/d3ta-go/system/system/migration/rdbms"
	"gorm.io/gorm"
)

// CasbinRule Optimation

// IamCasbinRule type
// Original: https://github.com/casbin/gorm-adapter/blob/master/adapter.go#L31
type IamCasbinRule struct {
	TablePrefix string `gorm:"-"`
	PType       string `gorm:"size:100;index;index:idx_unique,unique"`
	V0          string `gorm:"size:100;index;index:idx_unique,unique"`
	V1          string `gorm:"size:100;index;index:idx_unique,unique"`
	V2          string `gorm:"size:100;index;index:idx_unique,unique"`
	V3          string `gorm:"size:100;index;index:idx_unique,unique"`
	V4          string `gorm:"size:100;index;index:idx_unique,unique"`
	V5          string `gorm:"size:100;index;index:idx_unique,unique"`
}

// TableName func
func (c *IamCasbinRule) TableName() string {
	return "iam_casbin_rule"
}

// Seed20201119002InitCasbin type
type Seed20201119002InitCasbin struct {
	migRDBMS.BaseGormMigratorRunner
}

// NewSeed20201119002InitCasbin constructor
func NewSeed20201119002InitCasbin(h *handler.Handler) (migRDBMS.IGormMigratorRunner, error) {
	gmr := new(Seed20201119002InitCasbin)
	gmr.SetHandler(h)
	gmr.SetID("Seed20201119002InitCasbin")
	return gmr, nil
}

// GetID get Seed20201119002InitCasbin ID
func (dmr *Seed20201119002InitCasbin) GetID() string {
	return fmt.Sprintf("%T", dmr)
}

// Run run Seed20201119002InitCasbin
func (dmr *Seed20201119002InitCasbin) Run(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr.GetGorm().AutoMigrate(&IamCasbinRule{}); err != nil {
			return err
		}
		if err := dmr._seeds(); err != nil {
			return err
		}
	}
	return nil
}

// RollBack rollback Seed20201119002InitCasbin
func (dmr *Seed20201119002InitCasbin) RollBack(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr._unSeeds(); err != nil {
			return err
		}
	}
	return nil
}

var cPs = []IamCasbinRule{
	// role:system - ...

	// role:admin - ...

	// role:default - ...
}

var vGs = []IamCasbinRule{
	// group -> role (for flexibility)
	{PType: "g", V0: "group:system", V1: "role:system"},
	{PType: "g", V0: "group:admin", V1: "role:admin"},
	{PType: "g", V0: "group:default", V1: "role:default"},
}

func (dmr *Seed20201119002InitCasbin) _seeds() error {
	if dmr.GetGorm().Migrator().HasTable(&IamCasbinRule{}) {
		if len(cPs) > 0 {
			if err := dmr.GetGorm().Create(&cPs).Error; err != nil {
				return err
			}
		}

		for _, v := range vGs {
			var ett IamCasbinRule
			result := dmr.GetGorm().Unscoped().Where(v).First(&ett)
			if result.RowsAffected < 1 {
				if err := dmr.GetGorm().Create(&v).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (dmr *Seed20201119002InitCasbin) _unSeeds() error {
	if dmr.GetGorm().Migrator().HasTable(&IamCasbinRule{}) {

		for _, v := range cPs {
			if err := dmr.GetGorm().Unscoped().Where(&v).Delete(&IamCasbinRule{}).Error; err != nil {
				return err
			}
		}

		for _, v := range vGs {
			var ett IamCasbinRule
			result := dmr.GetGorm().Unscoped().Where(&IamCasbinRule{PType: "p", V0: v.V1}).First(&ett)
			if result.RowsAffected < 1 {
				if err := dmr.GetGorm().Where(&v).Delete(&IamCasbinRule{}).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}
