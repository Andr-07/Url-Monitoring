package monitor_log

import (
	"go-monitoring/internal/models"
	"time"
)

type MonitorLogDto struct {
	URLID     uint      
	Timestamp time.Time 
	Status    models.MonitorStatus   
	HTTPCode  int       
	Error     string   
}