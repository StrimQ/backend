CREATE TYPE "tier" AS ENUM (
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

CREATE TABLE "tenants" (
  "tenant_id" UUID,
  "name" varchar(255),
  "domain" varchar(255),
  "tier" tier,
  "infra_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id")
);

CREATE TABLE "users" (
  "user_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("user_id")
);

CREATE TABLE "tenant_users" (
  "tenant_id" UUID,
  "user_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "user_id")
);

CREATE TABLE "tenant_infras" (
  "tenant_infra_id" UUID,
  "name" varchar(255),
  "kafka_brokers" varchar(255)[],
  "schema_registry_url" varchar(255),
  "kafka_connect_url" varchar(255),
  "kms_key" varchar(255),
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_infra_id")
);

CREATE TABLE "tags" (
  "tenant_id" UUID,
  "tag_id" UUID,
  "key" varchar(255),
  "value" varchar(255),
  "created_by_user_id" UUID,
  "updated_by_user_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "tag_id")
);

CREATE TABLE "topics" (
  "tenant_id" UUID,
  "topic_id" UUID,
  "name" varchar(255),
  "producer_type" topic_producer_type,
  "producer_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "topic_id")
);

CREATE TABLE "sources" (
  "tenant_id" UUID,
  "source_id" UUID,
  "name" varchar(255),
  "engine" source_engine,
  "config" JSONB,
  "created_by_user_id" UUID,
  "updated_by_user_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "source_id")
);

CREATE TABLE "source_outputs" (
  "tenant_id" UUID,
  "source_id" UUID,
  "topic_id" UUID,
  "database_name" VARCHAR(255),
  "group_name" VARCHAR(255),
  "collection_name" VARCHAR(255),
  "config" JSONB,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "source_id", "topic_id")
);

CREATE TABLE "transformers" (
  "tenant_id" UUID,
  "transfomer_id" UUID,
  "name" varchar(255),
  "config" JSONB,
  "created_by_user_id" UUID,
  "updated_by_user_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "transfomer_id")
);

CREATE TABLE "transformer_inputs" (
  "tenant_id" UUID,
  "transformer_id" UUID,
  "topic_id" UUID,
  "config" JSONB,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "transformer_id", "topic_id")
);

CREATE TABLE "transformer_outputs" (
  "tenant_id" UUID,
  "transformer_id" UUID,
  "topic_id" UUID,
  "config" JSONB,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "transformer_id", "topic_id")
);

CREATE TABLE "destinations" (
  "tenant_id" UUID,
  "destination_id" UUID,
  "name" varchar(255),
  "engine" destination_engine,
  "config" JSONB,
  "created_by_user_id" UUID,
  "updated_by_user_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "destination_id")
);

CREATE TABLE "pipelines" (
  "tenant_id" UUID,
  "pipeline_id" UUID,
  "name" varchar(255),
  "source_id" UUID,
  "destination_id" UUID,
  "config" JSONB,
  "created_by_user_id" UUID,
  "updated_by_user_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "pipeline_id")
);

CREATE TABLE "pipeline_transformers" (
  "tenant_id" UUID,
  "pipeline_id" UUID,
  "transformer_id" UUID,
  "stage" int,
  "config" JSONB,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "pipeline_id", "transformer_id")
);

CREATE TABLE "pipeline_transformer_inputs" (
  "tenant_id" UUID,
  "pipeline_id" UUID,
  "transformer_id" UUID,
  "topic_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "pipeline_id", "transformer_id", "topic_id")
);

CREATE TABLE "pipeline_destination_inputs" (
  "tenant_id" UUID,
  "pipeline_id" UUID,
  "topic_id" UUID,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("tenant_id", "pipeline_id", "topic_id")
);

CREATE UNIQUE INDEX ON "pipeline_transformers" ("tenant_id", "pipeline_id", "stage");

ALTER TABLE "tenants" ADD FOREIGN KEY ("infra_id") REFERENCES "tenant_infras" ("tenant_infra_id");

ALTER TABLE "tenant_users" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("tenant_id");

ALTER TABLE "tenant_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "tags" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("tenant_id");

ALTER TABLE "tags" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "tags" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "topics" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("tenant_id");

ALTER TABLE "sources" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("tenant_id");

ALTER TABLE "sources" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "sources" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "source_outputs" ADD FOREIGN KEY ("tenant_id", "source_id") REFERENCES "sources" ("tenant_id", "source_id");

ALTER TABLE "source_outputs" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topics" ("tenant_id", "topic_id");

ALTER TABLE "transformers" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("tenant_id");

ALTER TABLE "transformers" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "transformers" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "transformer_inputs" ADD FOREIGN KEY ("tenant_id", "transformer_id") REFERENCES "transformers" ("tenant_id", "transfomer_id");

ALTER TABLE "transformer_inputs" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topics" ("tenant_id", "topic_id");

ALTER TABLE "transformer_outputs" ADD FOREIGN KEY ("tenant_id", "transformer_id") REFERENCES "transformers" ("tenant_id", "transfomer_id");

ALTER TABLE "transformer_outputs" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topics" ("tenant_id", "topic_id");

ALTER TABLE "destinations" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("tenant_id");

ALTER TABLE "destinations" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "destinations" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "pipelines" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("tenant_id");

ALTER TABLE "pipelines" ADD FOREIGN KEY ("tenant_id", "source_id") REFERENCES "sources" ("tenant_id", "source_id");

ALTER TABLE "pipelines" ADD FOREIGN KEY ("tenant_id", "destination_id") REFERENCES "destinations" ("tenant_id", "destination_id");

ALTER TABLE "pipelines" ADD FOREIGN KEY ("created_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "pipelines" ADD FOREIGN KEY ("updated_by_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "pipeline_transformers" ADD FOREIGN KEY ("tenant_id", "pipeline_id") REFERENCES "pipelines" ("tenant_id", "pipeline_id");

ALTER TABLE "pipeline_transformers" ADD FOREIGN KEY ("tenant_id", "transformer_id") REFERENCES "transformers" ("tenant_id", "transfomer_id");

ALTER TABLE "pipeline_transformer_inputs" ADD FOREIGN KEY ("tenant_id", "pipeline_id", "transformer_id") REFERENCES "pipeline_transformers" ("tenant_id", "pipeline_id", "transformer_id");

ALTER TABLE "pipeline_transformer_inputs" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topics" ("tenant_id", "topic_id");

ALTER TABLE "pipeline_destination_inputs" ADD FOREIGN KEY ("tenant_id", "pipeline_id") REFERENCES "pipelines" ("tenant_id", "pipeline_id");

ALTER TABLE "pipeline_destination_inputs" ADD FOREIGN KEY ("tenant_id", "topic_id") REFERENCES "topics" ("tenant_id", "topic_id");
