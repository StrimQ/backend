import uuid
from datetime import datetime
from enum import Enum as PyEnum

from sqlalchemy import Column, Date, Enum, ForeignKey, String, Table, func
from sqlalchemy.dialects.postgresql import ARRAY, UUID
from sqlalchemy.orm import Mapped, mapped_column, relationship

from app.models.base import Base


class Tier(PyEnum):
    FREE_TRIAL = "free_trial"
    BRONZE = "bronze"
    SILVER = "silver"
    GOLD = "gold"
    PLATINUM = "platinum"


users_tenants = Table(
    "users_tenants",
    Base.metadata,
    Column("user_id", ForeignKey("users.id"), primary_key=True),
    Column("tenant_id", ForeignKey("tenants.id"), primary_key=True),
    Column("created_at", Date, server_default=func.now()),
    Column(
        "updated_at",
        Date,
        server_default=func.now(),
        server_onupdate=func.now(),
    ),
)


class Tenant(Base):
    __tablename__ = "tenants"

    id: Mapped[uuid.UUID] = mapped_column(UUID(as_uuid=True), primary_key=True)
    name: Mapped[str] = mapped_column(String(255))
    domain: Mapped[str] = mapped_column(String(255))
    tier: Mapped[Tier] = mapped_column(Enum(Tier))
    infra_id: Mapped[uuid.UUID] = mapped_column(ForeignKey("tenant_infras.id"))
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )

    infra: Mapped["TenantInfra"] = relationship(back_populates="tenants")
    users: Mapped[list["User"]] = relationship(secondary=users_tenants)


class User(Base):
    __tablename__ = "users"

    id: Mapped[uuid.UUID] = mapped_column(UUID(as_uuid=True), primary_key=True)
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )


class TenantInfra(Base):
    __tablename__ = "tenant_infras"

    id: Mapped[uuid.UUID] = mapped_column(UUID(as_uuid=True), primary_key=True)
    name: Mapped[str] = mapped_column(String(255))
    kafka_brokers: Mapped[list[str]] = mapped_column(ARRAY(String(255)))
    schema_registry_url: Mapped[str] = mapped_column(String(255))
    kafka_connect_url: Mapped[str] = mapped_column(String(255))
    kms_key: Mapped[str] = mapped_column(String(255))
    created_at: Mapped[datetime] = mapped_column(server_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        server_default=func.now(), server_onupdate=func.now()
    )

    tenants: Mapped[list["Tenant"]] = relationship(back_populates="infra")
