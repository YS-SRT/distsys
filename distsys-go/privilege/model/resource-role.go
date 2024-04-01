package model

import "gorm.io/gorm"

type ResourceRole struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	Lable       string `json:"label" gorm:"column:label;size:30;unique"`
	Description string `json:"description" gorm:"column:description;size:255"`

	IsActive  bool `json:"-" gorm:"column:is_active"`
	IsDeleted bool `json:"-" gorm:"column:is_deleted"`
}

func (ResourceRole) TableName() string {
	return "privilege_resource_roles"
}

type ResourceRoleMap struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	Resource string `json:"resource" gorm:"column:resource"`
	Action   string `json:"action" gorm:"column:action"`
	RoleId   int    `json:"role_id" gorm:"column:role_id"`

	IsActive  bool `json:"-" gorm:"column:is_active"`
	IsDeleted bool `json:"-" gorm:"column:is_deleted"`
}

func (ResourceRoleMap) TableName() string {
	return "privilege_resource_role_map"
}
