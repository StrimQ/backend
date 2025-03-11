from sqlalchemy.ext.asyncio import AsyncSession

from app.models.source import Source


class SourceRepository:
    def create_source(self, db: AsyncSession, source: Source) -> Source:
        pass


# Single instance created at module level (similar to @Repository)
source_repository = SourceRepository()
