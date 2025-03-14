from typing import Literal, Self

from pydantic import BaseModel, model_validator

from app.models.source import SourceAppConfig
from app.schemas.source.types import SQLSourceBinaryHandlingMode


class PostgreSQLSourceCreateConfig(BaseModel):
    connector: Literal["postgresql"] = "postgresql"
    host: str
    port: int = 5432
    username: str
    password: str
    database: str
    snapshot_table_schema: str = "public"
    slot_name: str
    publication_name: str
    binary_handling_mode: SQLSourceBinaryHandlingMode = (
        SQLSourceBinaryHandlingMode.BYTES
    )
    heartbeat_enabled: bool = False
    heartbeat_interval: int | None = None
    heartbeat_schema: str | None = None
    heartbeat_table: str | None = None

    # {
    #     "schema1": {
    #         "table1": ["column1", "column2"],
    #         "table2": ["column1", "column2"],
    #     },
    #     "schema2": {
    #         "table1": ["column1", "column2"],
    #         "table2": ["column1", "column2"],
    #     },
    # }
    table_hierarchy: dict[str, dict[str, list[str]]]

    @model_validator(mode="after")
    def validate_heartbeat_fields(self) -> Self:
        if self.heartbeat_enabled:
            if self.heartbeat_interval is None:
                raise ValueError(
                    "Heartbeat interval is required when heartbeat is enabled"
                )
            if self.heartbeat_schema is None:
                raise ValueError(
                    "Heartbeat schema is required when heartbeat is enabled"
                )
            if self.heartbeat_table is None:
                raise ValueError(
                    "Heartbeat table is required when heartbeat is enabled"
                )
        return self


    def to_app_configs_model(self) -> list[SourceAppConfig]:
        app_configs = [
        ]

        for k, v in [
            ("connector", self.connector),
            ("host", self.host),
            ("port", str(self.port)),
            ("username", self.username),
            ("password", self.password),
            ("database", self.database),
            ("heartbeat_enabled", str(self.heartbeat_enabled)),
            ("heartbeat_interval", str(self.heartbeat_interval)),
            ("heartbeat_schema", self.heartbeat_schema),
            ("heartbeat_table", self.heartbeat_table),
        ]:
            app_configs.append(SourceAppConfig(key=k, value=v))
