package middleware

import (
	"context"
	"net/http"

	"github.com/StrimQ/backend/internal/client"
	"github.com/StrimQ/backend/internal/domain"
	"github.com/StrimQ/backend/internal/enum"
	"github.com/google/uuid"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := domain.NewUser(uuid.Max, uuid.Max)

		ctx := context.WithValue(r.Context(), enum.ContextKey_User, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func InitKCClient(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(enum.ContextKey_User).(*domain.User)
		if !ok {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}
		kcClient := client.NewKafkaConnectClient(user.Tenant.Infra.KafkaConnectURL, 10)
		ctx := context.WithValue(r.Context(), enum.ContextKey_KCClient, kcClient)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
