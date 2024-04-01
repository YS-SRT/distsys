package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	Nickname   string `json:"nickname" gorm:"column:nickname;size:30"`
	WechatUIDS string `json:"wechat_uids,omitempty" gorm:"column:wechat_uids;size:255;uniqueIndex"`
	Portrait   string `json:"portrait,omitempty" gorm:"column:portrait;size:100"`
	Phone      string `json:"phone,omitempty" gorm:"column:phone;size:20;uniqueIndex"`
	Password   string `json:"-" gorm:"column:password;size:150"`

	Age int `json:"age,omitempty" gorm:"column:age"`

	AddrID  int `json:"addr_id,omitempty" gorm:"column:addr_id"`
	LevelID int `json:"level_id,omitempty" gorm:"column:level_id"`

	IsActive  bool `json:"-" gorm:"column:is_active"`
	IsDeleted bool `json:"-" gorm:"column:is_deleted"`

	Score int `json:"score,omitempty"  gorm:"column:score"`
}

func (User) TableName() string {
	return "user_info"
}
