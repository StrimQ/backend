import uuid

from sqlalchemy import Column, DateTime, ForeignKey, String, Text
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import DeclarativeBase, relationship
from utils.datetime_utils import aware_utcnow


class Destination(DeclarativeBase):
    __tablename__ = "destinations"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    tenant_id = Column(UUID(as_uuid=True), ForeignKey("tenants.id"), nullable=False)
    name = Column(String(255), nullable=False)
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    tenant = relationship("Tenant", back_populates="destinations")
    tags = relationship("DestinationTag", back_populates="destination")
    app_configs = relationship("DestinationAppConfig", back_populates="destination")
    kc_connectors = relationship("DestinationKCConnector", back_populates="destination")
    pipelines = relationship("Pipeline", back_populates="destination")


# Destination Tag
class DestinationTag(DeclarativeBase):
    __tablename__ = "destination_tags"

    destination_id = Column(
        UUID(as_uuid=True), ForeignKey("destinations.id"), primary_key=True
    )
    key = Column(String(255), primary_key=True)
    value = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    destination = relationship("Destination", back_populates="tags")


# Destination App Config
class DestinationAppConfig(DeclarativeBase):
    __tablename__ = "destination_app_configs"

    destination_id = Column(
        UUID(as_uuid=True), ForeignKey("destinations.id"), primary_key=True
    )
    key = Column(String(255), primary_key=True)
    value = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    destination = relationship("Destination", back_populates="app_configs")


# Destination Kafka Connect Connector
class DestinationKCConnector(DeclarativeBase):
    __tablename__ = "destination_kc_connectors"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    destination_id = Column(
        UUID(as_uuid=True), ForeignKey("destinations.id"), nullable=False
    )
    name = Column(String(255), nullable=False, unique=True)
    connector_class = Column(String(255))
    version = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    destination = relationship("Destination", back_populates="kc_connectors")
    configs = relationship("DestinationKCConfig", back_populates="kc_connector")


# Destination Kafka Connect Config
class DestinationKCConfig(DeclarativeBase):
    __tablename__ = "destination_kc_configs"

    kc_id = Column(
        UUID(as_uuid=True), ForeignKey("destination_kc_connectors.id"), primary_key=True
    )
    key = Column(String(255), primary_key=True)
    value = Column(Text)
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    kc_connector = relationship("DestinationKCConnector", back_populates="configs")
