package model

import "gorm.io/gorm"

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:200"`
	V0    string `gorm:"size:200"`
	V1    string `gorm:"size:200"`
	V2    string `gorm:"size:200"`
	V3    string `gorm:"size:200"`
	V4    string `gorm:"size:200"`
	V5    string `gorm:"size:200"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}

type RestAction struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	Lable       string `json:"label" gorm:"column:label;size:30;unique"`
	Description string `json:"description" gorm:"column:description;size:200"`

	IsActive  bool `json:"-" gorm:"column:is_active"`
	IsDeleted bool `json:"-" gorm:"column:is_deleted"`
}

func (RestAction) TableName() string {
	return "privilege_rest_action"
}

type UserResourceGroupMap struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	UserRoleID        int    `json:"user_role_id" gorm:"user_role_id"`
	UserRoleLable     string `json:"user_role_label" gorm:"column:user_role_label;size:30"`
	ResourceRoleID    int    `json:"resource_role_id" gorm:"column:rescource_role_id"`
	ResourceRoleLable string `json:"resource_role_label" gorm:"column:resource_role_label;size:30"`

	ActionID    int    `json:"action_id" gorm:"action_id"`
	ActionLable string `json:"action_label" gorm:"column:action_label;size:30"`

	IsActive  bool `json:"-" gorm:"column:is_active"`
	IsDeleted bool `json:"-" gorm:"column:is_deleted"`
}

func (UserResourceGroupMap) TableName() string {
	return "privilege_user_resorce_role_map"
}
