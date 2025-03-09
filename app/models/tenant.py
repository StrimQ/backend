import enum

from sqlalchemy import Column, DateTime, Enum, ForeignKey, String
from sqlalchemy.dialects.postgresql import ARRAY, UUID
from sqlalchemy.orm import DeclarativeBase, relationship
from utils.datetime_utils import aware_utcnow


class Tier(enum.Enum):
    FREE_TRIAL = "free_trial"
    BRONZE = "bronze"
    SILVER = "silver"
    GOLD = "gold"
    PLATINUM = "platinum"


class Tenant(DeclarativeBase):
    __tablename__ = "tenants"

    id = Column(UUID(as_uuid=True), primary_key=True)
    name = Column(String(255), nullable=False)
    domain = Column(String(255), nullable=False)
    tier = Column(Enum(Tier), nullable=False)
    infra_id = Column(UUID(as_uuid=True), ForeignKey("tenant_infras.id"))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())
    deleted_at = Column(DateTime, nullable=True, default=None)

    infra = relationship("TenantInfra", back_populates="tenants")
    sources = relationship("Source", back_populates="tenant")
    destinations = relationship("Destination", back_populates="tenant")
    pipelines = relationship("Pipeline", back_populates="tenant")


class User(DeclarativeBase):
    __tablename__ = "users"

    id = Column(UUID(as_uuid=True), primary_key=True)
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())
    deleted_at = Column(DateTime, nullable=True, default=None)

    tenant = relationship("Tenant", back_populates="users")


class TenantInfra(DeclarativeBase):
    __tablename__ = "tenant_infras"

    id = Column(UUID(as_uuid=True), primary_key=True)
    name = Column(String(255), nullable=False)
    kafka_brokers = Column(ARRAY(String(255)))
    schema_registry_url = Column(String(255))
    kafka_connect_url = Column(String(255))
    kms_key = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())
    deleted_at = Column(DateTime, nullable=True, default=None)

    tenants = relationship("Tenant", back_populates="infra")
