package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	Lable       string `json:"label" gorm:"column:label;size:30;unique"`
	Description string `json:"description" gorm:"column:description;size:255"`

	IsActive  bool `json:"-" gorm:"column:is_active"`
	IsDeleted bool `json:"-" gorm:"column:is_deleted"`
}

func (UserRole) TableName() string {
	return "privilege_user_roles"
}

type UserRoleMap struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	UserID int `json:"uid" gorm:"column:user_id"`
	RoleId int `json:"role_id" gorm:"column:role_id"`

	IsActive  bool `json:"-" gorm:"column:is_active"`
	IsDeleted bool `json:"-" gorm:"column:is_deleted"`
}

func (UserRoleMap) TableName() string {
	return "privilege_user_role_map"
}
