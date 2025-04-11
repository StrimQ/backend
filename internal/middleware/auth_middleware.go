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
		// TODO: Implement actual authentication logic
		user := domain.NewUser(uuid.Max, uuid.Max)
		user.Tenant = &domain.Tenant{
			TenantID: uuid.Max,
			Infra: &domain.TenantInfra{
				KafkaBrokers:      []string{"kafka-0:9092", "kafka-1:9092"},
				SchemaRegistryURL: "http://localhost:8081",
				KafkaConnectURL:   "http://localhost:8083",
				KmsKey:            "arn:aws:kms:ap-southeast-1:123456789012:key/ffffffff-ffff-ffff-ffff-ffffffffffff",
			},
		}

		ctx := context.WithValue(r.Context(), enum.ContextKey_User, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func InjectKCClient(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(enum.ContextKey_User).(*domain.User)
		if !ok {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}
		kcClient := client.NewKafkaConnectClient(user.Tenant.Infra.KafkaConnectURL)
		ctx := context.WithValue(r.Context(), enum.ContextKey_KCClient, kcClient)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
