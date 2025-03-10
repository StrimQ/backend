from sqlalchemy import URL
from sqlalchemy.ext.asyncio import async_sessionmaker, create_async_engine

from app.settings import settings

# PostgreSQL engine
pg_conn_str = URL.create(
    drivername="postgresql+asyncpg",
    username=settings.PG_DB.USERNAME,
    password=settings.PG_DB.PASSWORD,
    host=settings.PG_DB.HOST,
    port=settings.PG_DB.PORT,
    database=settings.PG_DB.DBNAME,
)
pg_engine = create_async_engine(pg_conn_str)
pg_sessionmaker = async_sessionmaker(
    bind=pg_engine,
    expire_on_commit=False,
    autoflush=False,
)


# Dependency injection for FastAPI
async def get_pg_db():
    async with pg_sessionmaker() as pg_session:
        try:
            yield pg_session
            await pg_session.commit()
        except Exception:
            await pg_session.rollback()
            raise
