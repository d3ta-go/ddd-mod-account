package entity

// SysUserClientAppsEntity represent SysUserClientApps Entity
type SysUserClientAppsEntity struct {
	ID uint64 `gorm:"primary_key;column:id"`

	UUID          string `json:"uuid" gorm:"column:uuid;size:255;unique;not null"`
	ClientAppCode string `json:"clientAppCode" gorm:"column:client_app_code;size:255;not null;unique"`
	ClientAppName string `json:"clientAppName" gorm:"column:client_app_name;size:255;not null"`
	ClientAppDesc string `json:"ClientAppDesc" gorm:"column:client_app_desc;size:255;"`
	ClientKey     string `json:"clientKey" gorm:"column:client_key;size:255;"`
	SecretKey     string `json:"secretKey" gorm:"column:secret_key;size:500;"`
	IsActive      bool   `json:"isActive" gorm:"column:is_active;index"`

	User   SysUserEntity `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:ID;"`
	UserID uint64        `json:"userID" gorm:"column:user_id;"`

	BaseEntity
}

// TableName get real database table name
func (t *SysUserClientAppsEntity) TableName() string {
	return "iam_sys_user_client_apps"
}
