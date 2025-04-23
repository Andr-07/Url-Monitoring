package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Address  string `gorm:"uniqueIndex;not null" json:"address"`
	Interval int    `gorm:"default:60" json:"interval"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
