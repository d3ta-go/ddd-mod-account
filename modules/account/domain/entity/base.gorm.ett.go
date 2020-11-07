package entity

import (
	"time"

	"gorm.io/gorm"
)

// BaseEntity represent Base Entity
type BaseEntity struct {
	CreatedBy string         `json:"createdBy,omitempty" gorm:"column:sys_created_by;size:255" `
	CreatedAt *time.Time     `json:"createdAt,omitempty" gorm:"column:sys_created_at"`
	UpdatedBy string         `json:"updatedBy,omitempty" gorm:"column:sys_updated_by;size:255"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty" gorm:"column:sys_updated_at"`
	DeletedBy string         `json:"deletedBy,omitempty" gorm:"column:sys_deleted_by;size:255"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;index:idx_delete_at"`
	// DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"column:sys_deleted_at"`
}
