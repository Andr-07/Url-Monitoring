package url

import (
	"fmt"
	"go-monitoring/config"
	"go-monitoring/internal/middleware"
	"go-monitoring/pkg/req"
	"go-monitoring/pkg/res"
	"net/http"
)

type UrlHandlerDeps struct {
	*configs.Config
	*UrlService
}

type UrlHandler struct {
	*configs.Config
	*UrlService
}

func NewUrlHandler(router *http.ServeMux, deps UrlHandlerDeps) {
	handler := &UrlHandler{
		Config:     deps.Config,
		UrlService: deps.UrlService,
	}
	router.Handle("POST /url", middleware.IsAuthed(handler.Create(), deps.Config))

}

func (handler *UrlHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.ContextUserKey).(uint)
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		address, err := handler.UrlService.Create(userId, body.Address, body.Interval)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := PostResponse{
			Message: fmt.Sprintf("address %s was added successfully", address),
		}
		res.Json(w, data, http.StatusAccepted)
	}
}
