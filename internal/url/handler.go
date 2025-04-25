package url

import (
	"fmt"
	"go-monitoring/config"
	"go-monitoring/internal/middleware"
	"go-monitoring/pkg/req"
	"go-monitoring/pkg/res"
	"net/http"
	"strconv"
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
	router.Handle("GET /url", middleware.IsAuthed(handler.GetAll(), deps.Config))
	router.Handle("POST /url", middleware.IsAuthed(handler.Create(), deps.Config))
	router.Handle("DELETE /url/{id}", middleware.IsAuthed(handler.Delete(), deps.Config))

}

func (handler *UrlHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.ContextUserKey).(uint)
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		address, err := handler.UrlService.Create(userId, body.Address)
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

func (handler *UrlHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.ContextUserKey).(uint)
		idString := r.PathValue("id")
		parsedID, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		id := uint(parsedID)
		err = handler.UrlService.Delete(id, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := PostResponse{
			Message: fmt.Sprintf("address id %d was deleted successfully", id),
		}
		res.Json(w, data, http.StatusAccepted)
	}
}

func (handler *UrlHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.ContextUserKey).(uint)
		urls, err := handler.UrlService.GetAll(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, urls, http.StatusAccepted)
	}
}
