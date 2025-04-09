CREATE TYPE "tenant_tier" AS ENUM (
  'free_trial',
  'bronze',
  'silver',
  'gold',
  'platinum'
);

CREATE TYPE "topic_producer_type" AS ENUM (
  'source',
  'transformer'
);

CREATE TYPE "source_engine" AS ENUM (
  'mysql',
  'postgresql'
);

CREATE TYPE "destination_engine" AS ENUM (
  'mysql',
  'postgresql'
);

CREATE TABLE "tenant" (
  "tenant_id" UUID,
  "name" varchar(255) NOT NULL,
  "domain" varchar(255) NOT NULL,
  "tier" tenant_tier NOT NULL,
  "tenant_infra_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id")
);

CREATE TABLE "users" (
  "tenant_id" UUID,
  "user_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "user_id")
);

CREATE TABLE "tenant_infra" (
  "tenant_infra_id" UUID,
  "name" varchar(255) NOT NULL,
  "kafka_brokers" varchar(255)[] NOT NULL,
  "schema_registry_url" varchar(255) NOT NULL,
  "kafka_connect_url" varchar(255) NOT NULL,
  "kms_key" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_infra_id")
);

CREATE TABLE "tag" (
  "tenant_id" UUID,
  "tag_id" UUID,
  "key" varchar(255) NOT NULL,
  "value" varchar(255) NOT NULL,
  "created_by_user_id" UUID NOT NULL,
  "updated_by_user_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "tag_id")
);

CREATE TABLE "topic" (
  "tenant_id" UUID,
  "topic_id" UUID,
  "name" varchar(255) NOT NULL,
  "producer_type" topic_producer_type NOT NULL,
  "producer_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "topic_id")
);

CREATE TABLE "source" (
  "source_id" UUID,
  "tenant_id" UUID,
  "name" varchar(255) NOT NULL,
  "engine" source_engine NOT NULL,
  "config" JSONB NOT NULL,
  "created_by_user_id" UUID NOT NULL,
  "updated_by_user_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "source_id")
);

CREATE TABLE "source_tag" (
  "tenant_id" UUID,
  "source_id" UUID,
  "tag_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "source_id", "tag_id")
);

CREATE TABLE "source_collection" (
  "tenant_id" UUID,
  "source_id" UUID,
  "topic_id" UUID,
  "database_name" VARCHAR(255) NOT NULL,
  "group_name" VARCHAR(255) NOT NULL,
  "collection_name" VARCHAR(255) NOT NULL,
  "config" JSONB NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "source_id", "topic_id")
);

CREATE TABLE "transformer" (
  "tenant_id" UUID,
  "transformer_id" UUID,
  "name" varchar(255) NOT NULL,
  "config" JSONB NOT NULL,
  "created_by_user_id" UUID NOT NULL,
  "updated_by_user_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "transformer_id")
);

CREATE TABLE "transformer_tag" (
  "tenant_id" UUID,
  "transformer_id" UUID,
  "tag_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "transformer_id", "tag_id")
);

CREATE TABLE "transformer_input" (
  "tenant_id" UUID,
  "transformer_id" UUID,
  "topic_id" UUID,
  "config" JSONB NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "transformer_id", "topic_id")
);

CREATE TABLE "transformer_output" (
  "tenant_id" UUID,
  "transformer_id" UUID,
  "topic_id" UUID,
  "config" JSONB NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "transformer_id", "topic_id")
);

CREATE TABLE "destination" (
  "tenant_id" UUID,
  "destination_id" UUID,
  "name" varchar(255) NOT NULL,
  "engine" destination_engine NOT NULL,
  "config" JSONB NOT NULL,
  "created_by_user_id" UUID NOT NULL,
  "updated_by_user_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "destination_id")
);

CREATE TABLE "destination_collection" (
  "tenant_id" UUID,
  "destination_id" UUID,
  "topic_id" UUID,
  "database_name" VARCHAR(255) NOT NULL,
  "group_name" VARCHAR(255) NOT NULL,
  "collection_name" VARCHAR(255) NOT NULL,
  "config" JSONB NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "destination_id", "topic_id")
);

CREATE TABLE "destination_tag" (
  "tenant_id" UUID,
  "destination_id" UUID,
  "tag_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "destination_id", "tag_id")
);

CREATE TABLE "source_destination_stream" (
  "tenant_id" UUID,
  "source_id" UUID,
  "destination_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "source_id", "destination_id")
);

CREATE TABLE "source_transformer_stream" (
  "tenant_id" UUID,
  "source_id" UUID,
  "transformer_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "source_id", "transformer_id")
);

CREATE TABLE "transformer_recursive_stream" (
  "tenant_id" UUID,
  "head_transformer_id" UUID,
  "tail_transformer_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "head_transformer_id", "tail_transformer_id")
);

CREATE TABLE "transformer_destination_stream" (
  "tenant_id" UUID,
  "transformer_id" UUID,
  "destination_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "transformer_id", "destination_id")
);

ALTER TABLE "tenant" ADD FOREIGN KEY ("tenant_infra_id") REFERENCES "tenant_infra" ("tenant_infra_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "users" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "tag" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "tag" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "tag" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "topic" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_tag" ADD FOREIGN KEY ("tenant_id", "source_id") REFERENCES "source" ("tenant_id", "source_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_tag" ADD FOREIGN KEY ("tenant_id", "tag_id") REFERENCES "tag" ("tenant_id", "tag_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_collection" ADD FOREIGN KEY ("tenant_id", "source_id") REFERENCES "source" ("tenant_id", "source_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_collection" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topic" ("tenant_id", "topic_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_tag" ADD FOREIGN KEY ("tenant_id", "transformer_id") REFERENCES "transformer" ("tenant_id", "transformer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_tag" ADD FOREIGN KEY ("tenant_id", "tag_id") REFERENCES "tag" ("tenant_id", "tag_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_input" ADD FOREIGN KEY ("tenant_id", "transformer_id") REFERENCES "transformer" ("tenant_id", "transformer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_input" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topic" ("tenant_id", "topic_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_output" ADD FOREIGN KEY ("tenant_id", "transformer_id") REFERENCES "transformer" ("tenant_id", "transformer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_output" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topic" ("tenant_id", "topic_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination_collection" ADD FOREIGN KEY ("tenant_id", "destination_id") REFERENCES "destination" ("tenant_id", "destination_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination_collection" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topic" ("tenant_id", "topic_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination_tag" ADD FOREIGN KEY ("tenant_id", "destination_id") REFERENCES "destination" ("tenant_id", "destination_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination_tag" ADD FOREIGN KEY ("tenant_id", "tag_id") REFERENCES "tag" ("tenant_id", "tag_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_destination_stream" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_destination_stream" ADD FOREIGN KEY ("source_id") REFERENCES "source" ("source_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_destination_stream" ADD FOREIGN KEY ("destination_id") REFERENCES "destination" ("destination_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_transformer_stream" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_transformer_stream" ADD FOREIGN KEY ("source_id") REFERENCES "source" ("source_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_transformer_stream" ADD FOREIGN KEY ("transformer_id") REFERENCES "transformer" ("transformer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_recursive_stream" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_recursive_stream" ADD FOREIGN KEY ("head_transformer_id") REFERENCES "transformer" ("transformer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_recursive_stream" ADD FOREIGN KEY ("tail_transformer_id") REFERENCES "transformer" ("transformer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_destination_stream" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_destination_stream" ADD FOREIGN KEY ("transformer_id") REFERENCES "transformer" ("transformer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_destination_stream" ADD FOREIGN KEY ("destination_id") REFERENCES "destination" ("destination_id") DEFERRABLE INITIALLY DEFERRED;
