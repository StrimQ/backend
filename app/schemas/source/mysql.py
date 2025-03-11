from typing import Literal, Self

from pydantic import BaseModel, model_validator


class MySQLSourceCreateConfig(BaseModel):
    connector: Literal["mysql"] = "mysql"
    host: str
    port: int = 3306
    database: str
    username: str
    password: str
    db_connection_tz: str = "UTC"
    heartbeat_enabled: bool = False
    heartbeat_interval: int | None = None
    heartbeat_schema: str | None = None
    heartbeat_table: str | None = None

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
