package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primarykey" json:"id"`

	CountryID int `json:"country_id" gorm:"column:country_id"`
	//same as Province
	RegionID int    `json:"region_id" gorm:"column:region_id"`
	CityID   int    `json:"city_id" gorm:"column:city_id"`
	AreaID   int    `json:"area_id" gorm:"column:area_id"`
	Street   string `json:"street,omitempty" gorm:"column:street;size:100"`
	RawAddr  string `json:"raw_addr,omitempty" gorm:"column:raw_addr;size:255"`

	Latitude  float32 `json:"latitude" gorm:"column:latitude"`
	Longitude float32 `json:"longitude" gorm:"column:longitude"`
}

func (Address) TableName() string {
	return "user_address"
}
