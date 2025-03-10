import enum
from typing import Annotated, Literal, Self, Union

from pydantic import BaseModel, Field, model_validator

from app.schemas.common import EntityNameQuery

ENTITY = "source"


class SQLSourceBinaryHandlingMode(str, enum.Enum):
    BYTES = "bytes"
    BASE64 = "base64"
    BASE64_URL_SAFE = "base64-url-safe"
    HEX = "hex"


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


class MySQLSourceCreateConfig(BaseModel):
    connector: Literal["mysql"] = "mysql"
    host: str
    port: int = 3306
    database: str
    username: str
    password: str

    # {
    #     "database1": {
    #         "table1": ["column1", "column2"],
    #         "table2": ["column1", "column2"],
    #     },
    #     "database2": {
    #         "table1": ["column1", "column2"],
    #         "table2": ["column1", "column2"],
    #     },
    # }
    table_hierarchy: dict[str, dict[str, list[str]]]


class SourceCreate(BaseModel):
    name: str = EntityNameQuery(ENTITY)
    config: Annotated[
        Union[PostgreSQLSourceCreateConfig, MySQLSourceCreateConfig],
        Field(discriminator="connector"),
    ]


class SourceResponse(BaseModel):
    pass
