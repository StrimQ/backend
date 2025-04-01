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
  "infra_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id")
);

CREATE TABLE "users" (
  "user_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("user_id")
);

CREATE TABLE "tenant_user" (
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
  "tenant_id" UUID,
  "source_id" UUID,
  "name" varchar(255) NOT NULL,
  "engine" source_engine NOT NULL,
  "config" JSONB NOT NULL,
  "created_by_user_id" UUID NOT NULL,
  "updated_by_user_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "source_id")
);

CREATE TABLE "source_output" (
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
  "transfomer_id" UUID,
  "name" varchar(255) NOT NULL,
  "config" JSONB NOT NULL,
  "created_by_user_id" UUID NOT NULL,
  "updated_by_user_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "transfomer_id")
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

CREATE TABLE "pipeline" (
  "tenant_id" UUID,
  "pipeline_id" UUID,
  "name" varchar(255) NOT NULL,
  "source_id" UUID,
  "destination_id" UUID NOT NULL,
  "config" JSONB NOT NULL,
  "created_by_user_id" UUID NOT NULL,
  "updated_by_user_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "pipeline_id")
);

CREATE TABLE "pipeline_transformer" (
  "tenant_id" UUID,
  "pipeline_id" UUID,
  "transformer_id" UUID,
  "stage" int NOT NULL,
  "config" JSONB NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "pipeline_id", "transformer_id")
);

CREATE TABLE "pipeline_transformer_input" (
  "tenant_id" UUID,
  "pipeline_id" UUID,
  "transformer_id" UUID,
  "topic_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "pipeline_id", "transformer_id", "topic_id")
);

CREATE TABLE "pipeline_destination_input" (
  "tenant_id" UUID,
  "pipeline_id" UUID,
  "topic_id" UUID,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  PRIMARY KEY ("tenant_id", "pipeline_id", "topic_id")
);

CREATE UNIQUE INDEX ON "pipeline_transformer" ("tenant_id", "pipeline_id", "stage");

ALTER TABLE "tenant" ADD FOREIGN KEY ("infra_id") REFERENCES "tenant_infra" ("tenant_infra_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "tenant_user" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "tenant_user" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "tag" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "tag" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "tag" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "topic" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_output" ADD FOREIGN KEY ("tenant_id", "source_id") REFERENCES "source" ("tenant_id", "source_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "source_output" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topic" ("tenant_id", "topic_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_input" ADD FOREIGN KEY ("tenant_id", "transformer_id") REFERENCES "transformer" ("tenant_id", "transfomer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_input" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topic" ("tenant_id", "topic_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_output" ADD FOREIGN KEY ("tenant_id", "transformer_id") REFERENCES "transformer" ("tenant_id", "transfomer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "transformer_output" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topic" ("tenant_id", "topic_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "destination" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("tenant_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline" ADD FOREIGN KEY ("tenant_id", "source_id") REFERENCES "source" ("tenant_id", "source_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline" ADD FOREIGN KEY ("tenant_id", "destination_id") REFERENCES "destination" ("tenant_id", "destination_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline_transformer" ADD FOREIGN KEY ("tenant_id", "pipeline_id") REFERENCES "pipeline" ("tenant_id", "pipeline_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline_transformer" ADD FOREIGN KEY ("tenant_id", "transformer_id") REFERENCES "transformer" ("tenant_id", "transfomer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline_transformer_input" ADD FOREIGN KEY ("tenant_id", "pipeline_id", "transformer_id") REFERENCES "pipeline_transformer" ("tenant_id", "pipeline_id", "transformer_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline_transformer_input" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topic" ("tenant_id", "topic_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline_destination_input" ADD FOREIGN KEY ("tenant_id", "pipeline_id") REFERENCES "pipeline" ("tenant_id", "pipeline_id") DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE "pipeline_destination_input" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topic" ("tenant_id", "topic_id") DEFERRABLE INITIALLY DEFERRED;
