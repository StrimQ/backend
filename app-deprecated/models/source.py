import uuid
from datetime import datetime
from typing import TYPE_CHECKING

from sqlalchemy import ForeignKey, String, Text, UniqueConstraint, func, text
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import Mapped, mapped_column, relationship

from app.models.base import Base

if TYPE_CHECKING:
    from app.models.tenant import User


class Source(Base):
    __tablename__ = "sources"

    id: Mapped[uuid.UUID] = mapped_column(
        UUID(as_uuid=True), primary_key=True, server_default=text("GEN_RANDOM_UUID()")
    )
    tenant_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("tenants.id"))
    name: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )

    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
    tags: Mapped[list["SourceTag"]] = relationship()
    app_configs: Mapped[list["SourceAppConfig"]] = relationship()
    app_tables: Mapped[list["SourceAppTable"]] = relationship()
    kc_connectors: Mapped[list["SourceKcConnector"]] = relationship()


class SourceTag(Base):
    __tablename__ = "source_tags"

    source_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("sources.id"), primary_key=True
    )
    key: Mapped[str] = mapped_column(String(255), primary_key=True)
    value: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )

    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])


class SourceAppConfig(Base):
    __tablename__ = "source_app_configs"

    source_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("sources.id"), primary_key=True
    )
    key: Mapped[str] = mapped_column(String(255), primary_key=True)
    value: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )

    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])


class SourceAppTable(Base):
    __tablename__ = "source_app_tables"

    id: Mapped[uuid.UUID] = mapped_column(
        UUID(as_uuid=True), primary_key=True, server_default=text("GEN_RANDOM_UUID()")
    )
    source_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("sources.id"))
    db_name: Mapped[str] = mapped_column(String(255))
    schema_name: Mapped[str] = mapped_column(String(255))
    table_name: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )

    __table_args__ = (
        UniqueConstraint("source_id", "db_name", "schema_name", "table_name"),
    )

    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
    columns: Mapped[list["SourceAppColumn"]] = relationship()


class SourceAppColumn(Base):
    __tablename__ = "source_app_columns"

    table_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("source_app_tables.id"), primary_key=True
    )
    column_name: Mapped[str] = mapped_column(String(255), primary_key=True)
    data_type: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )

    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])


class SourceKcConnector(Base):
    __tablename__ = "source_kc_connectors"

    id: Mapped[uuid.UUID] = mapped_column(
        UUID(as_uuid=True), primary_key=True, server_default=text("GEN_RANDOM_UUID()")
    )
    source_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("sources.id"))
    name: Mapped[str] = mapped_column(String(255), unique=True)
    connector_class: Mapped[str] = mapped_column(String(255))
    version: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )

    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
    configs: Mapped[list["SourceKcConfig"]] = relationship()


class SourceKcConfig(Base):
    __tablename__ = "source_kc_configs"

    kc_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("source_kc_connectors.id"), primary_key=True
    )
    key: Mapped[str] = mapped_column(String(255), primary_key=True)
    value: Mapped[str] = mapped_column(Text)
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )

    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
