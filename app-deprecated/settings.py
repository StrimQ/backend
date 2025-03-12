from pydantic import BaseModel
from pydantic_settings import BaseSettings, SettingsConfigDict


class PostgreSQLSettings(BaseModel):
    HOST: str
    PORT: int
    USERNAME: str
    PASSWORD: str
    DBNAME: str


class Settings(BaseSettings):
    ENVIRONMENT: str = "development"
    PG_DB: PostgreSQLSettings

    model_config = SettingsConfigDict(
        env_file=".env",
        env_nested_delimiter="__",
    )


settings = Settings()
