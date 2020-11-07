package entity

// SysUserEntity represent SysUser Entity
type SysUserEntity struct {
	ID uint64 `json:"ID" gorm:"primary_key;column:id"`

	UUID      string `json:"uuid" gorm:"column:uuid;size:255;unique;not null"`
	Username  string `json:"userName" gorm:"column:username;size:255;unique;not null"`
	Password  string `json:"-" gorm:"column:password;size:255;not null;index"`
	NickName  string `json:"nickName" gorm:"column:nick_name;size:255;not null"`
	Email     string `json:"email" gorm:"column:email;size:255;unique;not null"`
	HeaderImg string `json:"headerImg" gorm:"column:header_img;size:255"`
	IsActive  bool   `json:"isActive" gorm:"column:is_active;index"`

	// Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId;"`
	AuthorityID string `json:"authorityId" gorm:"column:authority_id;size:255;not null;index"`

	BaseEntity
}

// TableName get real database table name
func (t *SysUserEntity) TableName() string {
	return "iam_sys_users"
}
