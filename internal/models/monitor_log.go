package models

import (
	"time"

	"gorm.io/gorm"
)

type MonitorStatus string

const (
	StatusOK   MonitorStatus = "OK"
	StatusFail MonitorStatus = "FAIL"
)

type MonitorLog struct {
	gorm.Model
	URLID     uint          `gorm:"not null" json:"url_id"`
	Timestamp time.Time     `gorm:"autoCreateTime" json:"timestamp"`
	Status    MonitorStatus `gorm:"type:varchar(10);not null" json:"status"` // "OK" / "FAIL"
	HTTPCode  int           `gorm:"not null" json:"http_code"`
	Error     string        `gorm:"type:text" json:"error"`
}
