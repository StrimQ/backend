package repository

import (
	"context"
	"encoding/json"

	"github.com/StrimQ/backend/internal/domain"
	"github.com/StrimQ/backend/internal/entity"
	"github.com/StrimQ/backend/internal/enum"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SourceRepository struct {
	db *pgxpool.Pool
}

func NewSourceRepository(db *pgxpool.Pool) *SourceRepository {
	return &SourceRepository{db}
}

func (r *SourceRepository) Create(ctx context.Context, source domain.Source) error {
	// Get the user from the context
	user := ctx.Value(domain.ContextKey_User).(*domain.User)
	userID := user.UserID

	// Begin a transaction
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	// Convert domain.Source to entity.SourceEntity
	metadata := source.GetMetadata()
	configJSON, err := json.Marshal(source.GetConfig())
	if err != nil {
		return err
	}

	sourceEntity := entity.SourceEntity{
		TenantID:        metadata.TenantID,
		SourceID:        metadata.SourceID,
		Name:            metadata.Name,
		Engine:          metadata.Engine,
		Config:          configJSON,
		CreatedByUserID: userID,
		UpdatedByUserID: userID,
	}

	// Insert SourceEntity
	_, err = tx.Exec(ctx, `
		INSERT INTO sources (tenant_id, source_id, name, engine, config, created_by_user_id, updated_by_user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, sourceEntity.TenantID, sourceEntity.SourceID, sourceEntity.Name, sourceEntity.Engine,
		sourceEntity.Config, sourceEntity.CreatedByUserID, sourceEntity.UpdatedByUserID)
	if err != nil {
		return err
	}

	// Generate and insert SourceOutputs
	outputs, err := source.DeriveOutputs()
	if err != nil {
		return err
	}

	for _, output := range outputs {
		topic, err := output.DeriveTopic()
		if err != nil {
			return err
		}

		topicEntity := entity.TopicEntity{
			TenantID:     sourceEntity.TenantID,
			TopicID:      topic.TopicID,
			Name:         topic.Name,
			ProducerType: enum.TopicProducerType_Source,
			ProducerID:   sourceEntity.SourceID,
		}

		_, err = tx.Exec(ctx, `
			INSERT INTO topics (tenant_id, topic_id, name, producer_type, producer_id)
			VALUES ($1, $2, $3, $4, $5)
		`, topicEntity.TenantID, topicEntity.TopicID, topicEntity.Name, topicEntity.ProducerType, topicEntity.ProducerID)
		if err != nil {
			return err
		}

		outputConfigJSON, err := json.Marshal(output.Config)
		if err != nil {
			return err
		}
		outputEntity := entity.SourceOutputEntity{
			TenantID:       sourceEntity.TenantID,
			SourceID:       sourceEntity.SourceID,
			TopicID:        topicEntity.TopicID,
			DatabaseName:   output.DatabaseName,
			GroupName:      output.GroupName,
			CollectionName: output.CollectionName,
			Config:         outputConfigJSON,
		}
		_, err = r.db.Exec(ctx, `
			INSERT INTO source_outputs (tenant_id, source_id, topic_id, database_name, group_name, collection_name, config)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, outputEntity.TenantID, outputEntity.SourceID, outputEntity.TopicID, outputEntity.DatabaseName,
			outputEntity.GroupName, outputEntity.CollectionName, outputEntity.Config)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	return tx.Commit(ctx)
}
