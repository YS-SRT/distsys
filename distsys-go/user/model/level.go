package model

import (
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	Lable    string `json:"label" gorm:"column:label;size:30"`
	MinScore int    `json:"min_score" gorm:"column:min_score"`
	MaxScore int    `json:"max_score" gorm:"column:max_score"`
}

func (Level) TableName() string {
	return "user_level"
}
