from uuid import UUID

from app.models.source import Source, SourceAppColumn, SourceAppConfig, SourceAppTable
from app.schemas.source.base import SourceCreate, SourceCreateConfig
from app.schemas.source.mysql import MySQLSourceCreateConfig
from app.schemas.source.postgresql import PostgreSQLSourceCreateConfig


class SourceMapper:
    @staticmethod
    def create_to_model(source_create: SourceCreate, tenant_id: UUID) -> Source:
        """
        Map a SourceCreate Pydantic model to a Source SQLAlchemy model with
            related entities.

        Args:
            source_create: The SourceCreate instance containing the source details.
            tenant_id: The UUID of the tenant owning this source.

        Returns:
            A fully constructed Source SQLAlchemy model instance.
        """
        # Create the base Source instance
        source = Source(name=source_create.name, tenant_id=tenant_id)
        config = source_create.config
        source.app_configs = []
        source.app_tables = []

        # Common configuration fields for both PostgreSQL and MySQL
        common_configs = [
            SourceAppConfig(key="connector", value=config.connector),
            SourceAppConfig(key="host", value=config.host),
            SourceAppConfig(key="port", value=str(config.port)),
            SourceAppConfig(key="username", value=config.username),
            SourceAppConfig(key="password", value=config.password),
            SourceAppConfig(key="database", value=config.database),
        ]
        source.app_configs.extend(common_configs)

        if isinstance(config, PostgreSQLSourceCreateConfig):
            # PostgreSQL-specific configurations
            pg_configs = [
                SourceAppConfig(
                    key="snapshot_table_schema", value=config.snapshot_table_schema
                ),
                SourceAppConfig(key="slot_name", value=config.slot_name),
                SourceAppConfig(key="publication_name", value=config.publication_name),
                SourceAppConfig(
                    key="binary_handling_mode",
                    value=config.binary_handling_mode.value,
                ),
                SourceAppConfig(
                    key="heartbeat_enabled", value=str(config.heartbeat_enabled)
                ),
            ]
            if config.heartbeat_enabled:
                pg_configs.extend(
                    [
                        SourceAppConfig(
                            key="heartbeat_interval",
                            value=str(config.heartbeat_interval),
                        ),
                        SourceAppConfig(
                            key="heartbeat_schema", value=config.heartbeat_schema
                        ),
                        SourceAppConfig(
                            key="heartbeat_table", value=config.heartbeat_table
                        ),
                    ]
                )
            source.app_configs.extend(pg_configs)

            # Map table_hierarchy for PostgreSQL (schemas within a single database)
            for schema, tables in config.table_hierarchy.items():
                for table, columns in tables.items():
                    app_table = SourceAppTable(
                        db_name=config.database, schema_name=schema, table_name=table
                    )
                    app_table.columns = [
                        SourceAppColumn(column_name=col) for col in columns
                    ]
                    source.app_tables.append(app_table)

        elif isinstance(config, MySQLSourceCreateConfig):
            # No additional MySQL-specific configs beyond the common ones

            # Map table_hierarchy for MySQL (multiple databases)
            for db, tables in config.table_hierarchy.items():
                for table, columns in tables.items():
                    app_table = SourceAppTable(
                        db_name=db,
                        schema_name="default",
                        table_name=table,
                    )
                    app_table.columns = [
                        SourceAppColumn(column_name=col) for col in columns
                    ]
                    source.app_tables.append(app_table)

        return source

    @staticmethod
    def create_config_to_model(
        source_create_config: SourceCreateConfig,
    ) -> list[SourceAppConfig]:
        pass
