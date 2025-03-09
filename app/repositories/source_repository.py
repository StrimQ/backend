from models import Source
from sqlalchemy.ext.asyncio import AsyncSession


class SourceRepository:
    def create_source(self, db: AsyncSession, name: str) -> Source:
        pass


# Single instance created at module level (similar to @Repository)
source_repository = SourceRepository()
