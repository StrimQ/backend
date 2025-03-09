import uuid

from sqlalchemy import Column, DateTime, ForeignKey, String, Text, UniqueConstraint
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import DeclarativeBase, relationship
from utils.datetime_utils import aware_utcnow


class Source(DeclarativeBase):
    __tablename__ = "sources"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    tenant_id = Column(UUID(as_uuid=True), ForeignKey("tenants.id"), nullable=False)
    name = Column(String(255), nullable=False)
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    tenant = relationship("Tenant", back_populates="sources")
    tags = relationship("SourceTag", back_populates="source")
    app_configs = relationship("SourceAppConfig", back_populates="source")
    app_tables = relationship("SourceAppTable", back_populates="source")
    kc_connectors = relationship("SourceKCConnector", back_populates="source")
    pipelines = relationship("Pipeline", back_populates="source")


class SourceTag(DeclarativeBase):
    __tablename__ = "source_tags"

    source_id = Column(UUID(as_uuid=True), ForeignKey("sources.id"), primary_key=True)
    key = Column(String(255), primary_key=True)
    value = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    source = relationship("Source", back_populates="tags")


class SourceAppTable(DeclarativeBase):
    __tablename__ = "source_app_tables"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    source_id = Column(UUID(as_uuid=True), ForeignKey("sources.id"), nullable=False)
    db_name = Column(String(255), nullable=False)
    schema_name = Column(String(255), nullable=False)
    table_name = Column(String(255), nullable=False)
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    __table_args__ = (
        UniqueConstraint(
            "source_id",
            "db_name",
            "schema_name",
            "table_name",
            name="uix_source_app_tables",
        ),
    )

    source = relationship("Source", back_populates="app_tables")
    columns = relationship("SourceAppColumn", back_populates="table")
    pipelines = relationship(
        "Pipeline",
        secondary="pipeline_source_app_tables",
        back_populates="source_app_tables",
    )


class SourceAppConfig(DeclarativeBase):
    __tablename__ = "source_app_configs"

    source_id = Column(UUID(as_uuid=True), ForeignKey("sources.id"), primary_key=True)
    key = Column(String(255), primary_key=True)
    value = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    source = relationship("Source", back_populates="app_configs")


class SourceAppColumn(DeclarativeBase):
    __tablename__ = "source_app_columns"

    table_id = Column(
        UUID(as_uuid=True), ForeignKey("source_app_tables.id"), primary_key=True
    )
    column_name = Column(String(255), primary_key=True)
    data_type = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    table = relationship("SourceAppTable", back_populates="columns")


# Source Kafka Connect Connector
class SourceKCConnector(DeclarativeBase):
    __tablename__ = "source_kc_connectors"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    source_id = Column(UUID(as_uuid=True), ForeignKey("sources.id"), nullable=False)
    name = Column(String(255), nullable=False, unique=True)
    connector_class = Column(String(255))
    version = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    source = relationship("Source", back_populates="kc_connectors")
    configs = relationship("SourceKCConfig", back_populates="kc_connector")


# Source Kafka Connect Config
class SourceKCConfig(DeclarativeBase):
    __tablename__ = "source_kc_configs"

    kc_id = Column(
        UUID(as_uuid=True), ForeignKey("source_kc_connectors.id"), primary_key=True
    )
    key = Column(String(255), primary_key=True)
    value = Column(Text)
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    kc_connector = relationship("SourceKCConnector", back_populates="configs")
