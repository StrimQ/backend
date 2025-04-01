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

// Create inserts a new source into the database along with its outputs and topics.
func (r *SourceRepository) Create(ctx context.Context, source *domain.Source) (*domain.Source, error) {
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

	// Generate and insert source outputs
	outputs, err := source.GenerateOutputs()
	if err != nil {
		return nil, fmt.Errorf("failed to generate source outputs: %w", err)
	}

	source.Outputs = outputs
	for _, output := range outputs {
		// Generate and insert the topic
		topic, err := output.GenerateTopic()
		if err != nil {
			return nil, fmt.Errorf("failed to generate topic for output: %w", err)
		}

		_, err = tx.Exec(ctx, `
			INSERT INTO topic (tenant_id, topic_id, name, producer_type, producer_id)
			VALUES ($1, $2, $3, $4, $5)
		`, source.TenantID, topic.TopicID, topic.Name, enum.TopicProducerType_Source, source.SourceID)
		if err != nil {
			return nil, fmt.Errorf("failed to insert topic: %w", err)
		}

		// Serialize the output config
		outputConfigJSON, err := json.Marshal(output.Config)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize output config: %w", err)
		}

		// Insert the source output
		_, err = tx.Exec(ctx, `
			INSERT INTO source_output (tenant_id, source_id, topic_id, database_name, group_name, collection_name, config)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, source.TenantID, source.SourceID, topic.TopicID, output.DatabaseName,
			output.GroupName, output.CollectionName, outputConfigJSON)
		if err != nil {
			return nil, fmt.Errorf("failed to insert source output: %w", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return source, nil
}
