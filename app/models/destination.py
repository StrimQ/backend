import uuid
from datetime import datetime
from typing import TYPE_CHECKING

from sqlalchemy import ForeignKey, String, Text, text
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import Mapped, mapped_column, relationship

from app.models.base import Base

if TYPE_CHECKING:
    from app.models.pipeline import Pipeline
    from app.models.tenant import Tenant, User


class Destination(Base):
    __tablename__ = "destinations"

    id: Mapped[uuid.UUID] = mapped_column(
        UUID(as_uuid=True), primary_key=True, server_default=text("gen_random_uuid()")
    )
    tenant_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("tenants.id"))
    name: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("now()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("now()"), server_onupdate=text("now()")
    )

    tenant: Mapped["Tenant"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
    tags: Mapped[list["DestinationTag"]] = relationship()
    app_configs: Mapped[list["DestinationAppConfig"]] = relationship()
    kc_connectors: Mapped[list["DestinationKcConnector"]] = relationship()
    pipelines: Mapped[list["Pipeline"]] = relationship()


class DestinationTag(Base):
    __tablename__ = "destination_tags"

    destination_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("destinations.id"), primary_key=True
    )
    key: Mapped[str] = mapped_column(String(255), primary_key=True)
    value: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("now()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("now()"), server_onupdate=text("now()")
    )

    destination: Mapped["Destination"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])


class DestinationAppConfig(Base):
    __tablename__ = "destination_app_configs"

    destination_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("destinations.id"), primary_key=True
    )
    key: Mapped[str] = mapped_column(String(255), primary_key=True)
    value: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("now()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("now()"), server_onupdate=text("now()")
    )

    destination: Mapped["Destination"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])


class DestinationKcConnector(Base):
    __tablename__ = "destination_kc_connectors"

    id: Mapped[uuid.UUID] = mapped_column(
        UUID(as_uuid=True), primary_key=True, server_default=text("gen_random_uuid()")
    )
    destination_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("destinations.id"))
    name: Mapped[str] = mapped_column(String(255), unique=True, nullable=False)
    connector_class: Mapped[str] = mapped_column(String(255))
    version: Mapped[str] = mapped_column(String(255))
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("now()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("now()"), server_onupdate=text("now()")
    )

    destination: Mapped["Destination"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
    configs: Mapped[list["DestinationKcConfig"]] = relationship()


class DestinationKcConfig(Base):
    __tablename__ = "destination_kc_configs"

    kc_id: Mapped[uuid.UUID] = mapped_column(
        ForeignKey("destination_kc_connectors.id"), primary_key=True
    )
    key: Mapped[str] = mapped_column(String(255), primary_key=True)
    value: Mapped[str] = mapped_column(Text)
    created_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    updated_by_user_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=text("now()"))
    updated_at: Mapped[datetime] = mapped_column(
        server_default=text("now()"), server_onupdate=text("now()")
    )

    connector: Mapped["DestinationKcConnector"] = relationship()
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_user_id])
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_user_id])
