package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/StrimQ/backend/internal/domain"
	"github.com/StrimQ/backend/internal/enum"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SourceRepository handles source-related database operations.
type SourceRepository struct {
	db *pgxpool.Pool
}

// NewSourceRepository creates a new SourceRepository instance.
func NewSourceRepository(db *pgxpool.Pool) *SourceRepository {
	return &SourceRepository{db}
}

// Add inserts a new source into the database along with its collections and topics.
func (r *SourceRepository) Add(ctx context.Context, source *domain.Source) (*domain.Source, error) {
	// Begin a transaction
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx) // Rollback if not committed
	}()

	// Serialize the source config to JSON
	configJSON, err := source.Config.AsBytes()
	if err != nil {
		return nil, fmt.Errorf("failed to serialize source config: %w", err)
	}

	// Insert the source into the database
	_, err = tx.Exec(ctx, `
		INSERT INTO source (tenant_id, source_id, name, engine, config, created_by_user_id, updated_by_user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, source.TenantID, source.SourceID, source.Name, source.Engine,
		configJSON, source.CreatedByUserID, source.UpdatedByUserID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert source: %w", err)
	}

	for _, collection := range source.Collections {
		// Generate and insert the topic
		topic := collection.Topic

		_, err = tx.Exec(ctx, `
			INSERT INTO topic (tenant_id, topic_id, name, producer_type, producer_id)
			VALUES ($1, $2, $3, $4, $5)
		`, source.TenantID, topic.TopicID, topic.Name, enum.TopicProducerType_Source, source.SourceID)
		if err != nil {
			return nil, fmt.Errorf("failed to insert topic: %w", err)
		}

		// Serialize the collection config
		collectionConfigJSON, err := json.Marshal(collection.Config)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize collection config: %w", err)
		}

		// Insert the source collection
		_, err = tx.Exec(ctx, `
			INSERT INTO source_collection (tenant_id, source_id, topic_id, database_name, group_name, collection_name, config)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, source.TenantID, source.SourceID, topic.TopicID, collection.DatabaseName,
			collection.GroupName, collection.CollectionName, collectionConfigJSON)
		if err != nil {
			return nil, fmt.Errorf("failed to insert source collection: %w", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return source, nil
}
