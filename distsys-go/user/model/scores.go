package model

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	UserID    int `json:"uid" gorm:"column:user_id"`
	ActID     int `json:"act_id" gorm:"column:activity_id"`
	ActTypeID int `json:"act_type_id" gorm:"column:act_type_id"`

	Score   int `json:"score" gorm:"column:score"`
	GroupId int `json:"group_id" gorm:"column:group_id"`

	FromUserID int `json:"from_uid" gorm:"column:from_user_id"`
	ToUserId   int `json:"to_uid" gorm:"column:to_user_id"`

	IsActive bool `json:"is_active" gorm:"column:is_active"`
	IsCalced bool `json:"is_calced" gorm:"column:is_calced"`
}

func (Score) TableName() string {
	return "user_score"
}
