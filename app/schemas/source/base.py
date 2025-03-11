from typing import Annotated

from pydantic import BaseModel, Field

from app.schemas.common import EntityNameQuery
from app.schemas.source.mysql import MySQLSourceCreateConfig
from app.schemas.source.postgresql import PostgreSQLSourceCreateConfig

ENTITY = "source"

SourceCreateConfig = Annotated[
    PostgreSQLSourceCreateConfig | MySQLSourceCreateConfig,
    Field(discriminator="connector"),
]


class SourceCreate(BaseModel):
    name: str = EntityNameQuery(ENTITY)
    config: SourceCreateConfig


class SourceResponse(BaseModel):
    pass
