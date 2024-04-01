package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	Lable       string `json:"label" gorm:"column:label;size:30"`
	Description string `json:"description" gorm:"column:description;size:255"`
	OwnerUID    int    `json:"owner_uid" gorm:"column:owner_uid"`

	IsActive  bool `json:"-" gorm:"column:is_active"`
	IsDeleted bool `json:"-" gorm:"column:is_deleted"`
}

func (Group) TableName() string {
	return "user_group"
}
