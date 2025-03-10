import uuid
from datetime import datetime
from typing import TYPE_CHECKING

from sqlalchemy import ForeignKey, String, text
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import Mapped, mapped_column, relationship

from app.models.base import Base

if TYPE_CHECKING:
    from app.models.destination import Destination
    from app.models.source import Source, SourceAppTable
    from app.models.tenant import Tenant, User


class Pipeline(Base):
    __tablename__ = "pipelines"

    id: Mapped[uuid.UUID] = mapped_column(
        UUID(as_uuid=True), primary_key=True, server_default=text("gen_random_uuid()")
    )
    tenant_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("tenants.id"))
    source_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("sources.id"))
    destination_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("destinations.id"))
    name: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("now()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("now()"), server_onupdate=text("now()")
    )

    tenant: Mapped["Tenant"] = relationship()
    source: Mapped["Source"] = relationship()
    destination: Mapped["Destination"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
    tags: Mapped[list["PipelineTag"]] = relationship()
    app_configs: Mapped[list["PipelineAppConfig"]] = relationship()
    source_tables: Mapped[list["PipelineSourceAppTable"]] = relationship()


class PipelineTag(Base):
    __tablename__ = "pipeline_tags"

    pipeline_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("pipelines.id"), primary_key=True
    )
    key: Mapped[str] = mapped_column(String(255), primary_key=True)
    value: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("now()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("now()"), server_onupdate=text("now()")
    )

    pipeline: Mapped["Pipeline"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])


class PipelineAppConfig(Base):
    __tablename__ = "pipeline_app_configs"

    pipeline_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("pipelines.id"), primary_key=True
    )
    key: Mapped[str] = mapped_column(String(255), primary_key=True)
    value: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("now()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("now()"), server_onupdate=text("now()")
    )

    pipeline: Mapped["Pipeline"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])


class PipelineSourceAppTable(Base):
    __tablename__ = "pipeline_source_app_tables"

    pipeline_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("pipelines.id"), primary_key=True
    )
    table_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("source_app_tables.id"), primary_key=True
    )
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("now()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("now()"), server_onupdate=text("now()")
    )

    pipeline: Mapped["Pipeline"] = relationship()
    table: Mapped["SourceAppTable"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
