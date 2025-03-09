from typing import Annotated, Literal, Self, Union

from pydantic import BaseModel, Field, model_validator

from app.schemas.common import EntityNameQuery

ENTITY = "source"


class SourceCreateBase(BaseModel):
    name: str = EntityNameQuery(ENTITY)


class PostgreSQLSourceCreate(SourceCreateBase):
    connector: Literal["postgresql"] = "postgresql"
    host: str
    port: int = 5432
    username: str
    password: str
    database: str
    snapshot_table_schema: str = "public"
    heartbeat_enabled: bool = False
    heartbeat_interval: int | None = None
    heartbeat_schema: str | None = None
    heartbeat_table: str | None = None

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


class MySQLSourceCreate(SourceCreateBase):
    connector: Literal["mysql"] = "mysql"
    host: str
    port: int = 3306
    database: str
    username: str
    password: str


SourceCreate = Annotated[
    Union[PostgreSQLSourceCreate, MySQLSourceCreate], Field(discriminator="connector")
]


class SourceResponse(BaseModel):
    pass
