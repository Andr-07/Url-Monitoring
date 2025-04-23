package middleware

import (
	"context"
	"go-monitoring/config"
	"go-monitoring/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextUserKey key = "ContextUserKey"
)

func writeUnathed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnathed(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnathed(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserKey, data.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
