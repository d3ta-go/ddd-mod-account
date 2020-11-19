package rdbms

import (
	"fmt"

	domEntity "github.com/d3ta-go/ddd-mod-account/modules/account/domain/entity"
	"github.com/d3ta-go/system/system/handler"
	migRDBMS "github.com/d3ta-go/system/system/migration/rdbms"
	"github.com/d3ta-go/system/system/utils"
	"gorm.io/gorm"
)

// Seed20201119001InitTable type
type Seed20201119001InitTable struct {
	migRDBMS.BaseGormMigratorRunner
}

// NewSeed20201119001InitTable constructor
func NewSeed20201119001InitTable(h *handler.Handler) (migRDBMS.IGormMigratorRunner, error) {
	gmr := new(Seed20201119001InitTable)
	gmr.SetHandler(h)
	gmr.SetID("Seed20201119001InitTable")
	return gmr, nil
}

// GetID get Seed20201119001InitTable ID
func (dmr *Seed20201119001InitTable) GetID() string {
	return fmt.Sprintf("%T", dmr)
}

// Run run Seed20201119001InitTable
func (dmr *Seed20201119001InitTable) Run(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr._seeds(); err != nil {
			return err
		}
	}
	return nil
}

// RollBack rollback Seed20201119001InitTable
func (dmr *Seed20201119001InitTable) RollBack(h *handler.Handler, dbGorm *gorm.DB) error {
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

func (dmr *Seed20201119001InitTable) _seeds() error {
	if dmr.GetGorm().Migrator().HasTable(&domEntity.SysUserEntity{}) &&
		dmr.GetGorm().Migrator().HasTable(&domEntity.SysUserClientAppsEntity{}) {
		// get default admin config
		cfg, err := dmr.GetHandler().GetDefaultConfig()
		if err != nil {
			return err
		}
		defaultAdmin := cfg.IAM.DefaultAdmin

		// create default user
		eUser := domEntity.SysUserEntity{
			UUID:        utils.GenerateUUID(),
			Username:    defaultAdmin.Username,
			Password:    utils.MD5([]byte(defaultAdmin.Password)),
			NickName:    defaultAdmin.NickName,
			Email:       defaultAdmin.Email,
			IsActive:    true,
			AuthorityID: defaultAdmin.AuthorityID,
		}
		eUser.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eUser).Error; err != nil {
			return err
		}

		// create default client app
		eCApp := domEntity.SysUserClientAppsEntity{
			UUID:          utils.GenerateUUID(),
			ClientAppCode: fmt.Sprintf("default-%s-app", defaultAdmin.Username),
			ClientAppName: fmt.Sprintf("Default %s App", defaultAdmin.NickName),
			ClientAppDesc: fmt.Sprintf("Default %s App Description", defaultAdmin.NickName),
			ClientKey:     utils.GenerateClientKey(),
			SecretKey:     utils.GenerateSecretKey(),
			IsActive:      true,
			UserID:        eUser.ID,
		}
		eCApp.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eCApp).Error; err != nil {
			return err
		}
	}
	return nil
}

func (dmr *Seed20201119001InitTable) _unSeeds() error {
	if dmr.GetGorm().Migrator().HasTable(&domEntity.SysUserEntity{}) &&
		dmr.GetGorm().Migrator().HasTable(&domEntity.SysUserClientAppsEntity{}) {
		// find default user
		cfg, err := dmr.GetHandler().GetDefaultConfig()
		if err != nil {
			return err
		}
		defaultAdmin := cfg.IAM.DefaultAdmin

		var eUser domEntity.SysUserEntity
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.SysUserEntity{Username: defaultAdmin.Username}).First(&eUser).Error; err != nil {
			return err
		}

		// delete default client app
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.SysUserClientAppsEntity{UserID: eUser.ID}).Delete(&domEntity.SysUserClientAppsEntity{}).Error; err != nil {
			return err
		}
		// delete default user
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.SysUserEntity{Username: eUser.Username}).Delete(&domEntity.SysUserEntity{}).Error; err != nil {
			return err
		}
	}
	return nil
}
