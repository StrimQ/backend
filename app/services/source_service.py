from repositories import source_repository
from schemas.source import SourceCreate, SourceResponse
from sqlalchemy.ext.asyncio import AsyncSession


class SourceService:

    def create_source(self, db: AsyncSession, source: SourceCreate) -> SourceResponse:
        source_repository.create_source(db, source)


# Single instance created at module level (similar to @Service)
source_service = SourceService()
