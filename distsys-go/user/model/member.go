package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	UserID      int `json:"uid" gorm:"column:user_id"`
	GroupId     int `json:"group_id" gorm:"column:group_id"`
	MemberLevel int `json:"member_level" gorm:"column:member_level"`

	IsCofounder bool `json:"is_cofounder" gorm:"column:is_cofounder"`
	IsDeleted   bool `json:"is_deleted" gorm:"column:is_deleted"`
}

func (Member) TableName() string {
	return "user_group_member"
}
