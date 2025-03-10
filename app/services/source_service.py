from schemas.source import SourceCreate, SourceResponse
from sqlalchemy.ext.asyncio import AsyncSession

from app.repositories.source_repository import source_repository


class SourceService:
    def create_source(self, db: AsyncSession, source: SourceCreate) -> SourceResponse:

        source_repository.create_source(db, source)


# Single instance created at module level (similar to @Service)
source_service = SourceService()
