package middleware

import (
	"context"
	"net/http"

	"github.com/StrimQ/backend/internal/domain"
	"github.com/google/uuid"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := domain.NewUser(uuid.Max, uuid.Max)

		ctx := context.WithValue(r.Context(), domain.ContextKey_User, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
