import uuid
from datetime import datetime
from typing import TYPE_CHECKING

from sqlalchemy import Column, DateTime, ForeignKey, String, Table, text
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import Mapped, mapped_column, relationship

from app.models.base import Base

if TYPE_CHECKING:
    from app.models.destination import Destination
    from app.models.source import Source, SourceAppTable
    from app.models.tenant import User


pipelines_source_app_tables = Table(
    "pipelines_source_app_tables",
    Base.metadata,
    Column("pipeline_id", ForeignKey("pipelines.id"), primary_key=True),
    Column("table_id", ForeignKey("source_app_tables.id"), primary_key=True),
    Column("created_at", DateTime, server_default=text("NOW()")),
    Column("created_by_user_id", UUID(as_uuid=True), ForeignKey("users.id")),
    Column("updated_by_user_id", UUID(as_uuid=True), ForeignKey("users.id")),
    Column(
        "updated_at",
        DateTime,
        server_default=text("NOW()"),
        server_onupdate=text("NOW()"),
    ),
)


class Pipeline(Base):
    __tablename__ = "pipelines"

    id: Mapped[uuid.UUID] = mapped_column(
        UUID(as_uuid=True), primary_key=True, server_default=text("GEN_RANDOM_UUID()")
    )
    tenant_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("tenants.id"))
    source_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("sources.id"))
    destination_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("destinations.id"))
    name: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("NOW()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("NOW()"), server_onupdate=text("NOW()")
    )

    source: Mapped["Source"] = relationship()
    destination: Mapped["Destination"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
    tags: Mapped[list["PipelineTag"]] = relationship()
    app_configs: Mapped[list["PipelineAppConfig"]] = relationship()
    source_tables: Mapped[list["SourceAppTable"]] = relationship(
        secondary=pipelines_source_app_tables
    )


class PipelineTag(Base):
    __tablename__ = "pipeline_tags"

    pipeline_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("pipelines.id"), primary_key=True
    )
    key: Mapped[str] = mapped_column(String(255), primary_key=True)
    value: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("NOW()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("NOW()"), server_onupdate=text("NOW()")
    )

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
    created_at: Mapped[datetime] = mapped_column(server_default=text("NOW()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("NOW()"), server_onupdate=text("NOW()")
    )

    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
