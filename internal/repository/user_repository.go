package repository

// import (
// 	"context"
// 	"fmt"

// 	"github.com/StrimQ/backend/internal/domain"
// 	"github.com/StrimQ/backend/internal/entity"
// 	"github.com/georgysavva/scany/v2/pgxscan"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type UserRepository struct {
// 	db *pgxpool.Pool
// }

// func NewUserRepository(db *pgxpool.Pool) *UserRepository {
// 	return &UserRepository{db}
// }

// // Get retrieves a user's complete information, including all associated tenants,
// // based on the UserID from the context (derived from JWT).
// func (r *UserRepository) Get(ctx context.Context) (domain.User, error) {
// 	// Extract user information from context (assumed to come from JWT)
// 	user, ok := ctx.Value(domain.ContextKey_User).(*domain.User)
// 	if !ok || user == nil {
// 		return domain.User{}, &UserNotFoundErr{}
// 	}
// 	userID := user.UserID

// 	// Begin a transaction
// 	tx, err := r.db.Begin(ctx)
// 	if err != nil {
// 		return domain.User{}, fmt.Errorf("failed to begin transaction: %w", err)
// 	}
// 	defer func() {
// 		_ = tx.Rollback(ctx)
// 	}()

// 	var userEntity entity.UserEntity
// 	pgxscan.Select(ctx, tx, &userEntity, `
// 		SELECT user_id, created_at, updated_at
// 		FROM users
// 		WHERE user_id = $1`, userID)

// 	// Step 2: Fetch all associated TenantEntity records via the tenant_user join table
// 	rows, err := r.db.Query(ctx, `
// 		SELECT t.tenant_id, t.name, t.domain, t.tier, t.infra_id, t.created_at, t.updated_at
// 		FROM tenant t
// 		JOIN tenant_user tu ON t.tenant_id = tu.tenant_id
// 		WHERE tu.user_id = $1
// 	`, userID)
// 	if err != nil {
// 		return domain.User{}, fmt.Errorf("failed to fetch tenants: %w", err)
// 	}
// 	defer rows.Close()

// 	// Collect tenant entities
// 	var tenants []entity.TenantEntity
// 	for rows.Next() {
// 		var tenant entity.TenantEntity
// 		err := rows.Scan(
// 			&tenant.TenantID,
// 			&tenant.Name,
// 			&tenant.Domain,
// 			&tenant.Tier,
// 			&tenant.InfraID,
// 			&tenant.CreatedAt,
// 			&tenant.UpdatedAt,
// 		)
// 		if err != nil {
// 			return domain.User{}, fmt.Errorf("failed to scan tenant: %w", err)
// 		}
// 		tenants = append(tenants, tenant)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return domain.User{}, fmt.Errorf("error iterating tenant rows: %w", err)
// 	}

// 	// Step 3: Map entities to domain.User
// 	// Assuming domain.User includes all UserEntity fields and a Tenants slice
// 	domainUser := domain.User{
// 		UserID:    userEntity.UserID,
// 		CreatedAt: userEntity.CreatedAt,
// 		UpdatedAt: userEntity.UpdatedAt,
// 		Tenants:   make([]domain.Tenant, len(tenants)),
// 	}

// 	// Map each TenantEntity to domain.Tenant
// 	for i, tenant := range tenants {
// 		domainUser.Tenants[i] = domain.Tenant{
// 			TenantID:  tenant.TenantID,
// 			Name:      tenant.Name,
// 			Domain:    tenant.Domain,
// 			Tier:      tenant.Tier,
// 			InfraID:   tenant.InfraID,
// 			CreatedAt: tenant.CreatedAt,
// 			UpdatedAt: tenant.UpdatedAt,
// 		}
// 	}

// 	return domainUser, nil
// }
