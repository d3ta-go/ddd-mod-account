package entity

import "time"

// TmpUserRegistrationEntity represent TmpUserRegistration Entity
type TmpUserRegistrationEntity struct {
	ID uint64 `json:"ID" gorm:"primary_key;column:id"`

	UUID     string `json:"uuid" gorm:"column:uuid;size:255;unique;not null"`
	Username string `json:"userName" gorm:"column:username;size:255;unique;not null"`
	Password string `json:"-" gorm:"column:password;size:255;not null"`
	NickName string `json:"nickName" gorm:"column:nick_name;size:255;not null"`
	Email    string `json:"email" gorm:"column:email;size:255;unique;not null;"`

	IsActivated    bool       `json:"isActivated" gorm:"column:is_activated;index"`
	ActivationCode string     `json:"activationCode" gorm:"column:activation_code;size:255"`
	ActivatedAt    *time.Time `json:"activateddAt,omitempty" gorm:"column:sys_activated_at"`

	BaseEntity
}

// TableName get real database table name
func (t *TmpUserRegistrationEntity) TableName() string {
	return "iam_tmp_user_registrations"
}
