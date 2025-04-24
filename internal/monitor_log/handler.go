package monitor_log

import (
	"go-monitoring/config"
	"go-monitoring/internal/middleware"
	"go-monitoring/pkg/res"
	"net/http"
	"strconv"
)

type MonitorLogHandlerDeps struct {
	*configs.Config
	*MonitorLogService
}

type MonitorLogHandler struct {
	*configs.Config
	*MonitorLogService
}

func NewMonitorLogHandler(router *http.ServeMux, deps MonitorLogHandlerDeps) {
	handler := &MonitorLogHandler{
		Config:            deps.Config,
		MonitorLogService: deps.MonitorLogService,
	}
	router.Handle("GET /url/{id}/logs", middleware.IsAuthed(handler.GetAll(), deps.Config))
}

func (handler *MonitorLogHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		parsedID, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		urlId := uint(parsedID)
		monitorLogs, err := handler.MonitorLogService.GetAll(urlId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, monitorLogs, http.StatusAccepted)
	}
}
