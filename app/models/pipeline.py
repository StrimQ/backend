import uuid

from sqlalchemy import Column, DateTime, ForeignKey, String
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import DeclarativeBase, relationship
from utils.datetime_utils import aware_utcnow


class Pipeline(DeclarativeBase):
    __tablename__ = "pipelines"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    tenant_id = Column(UUID(as_uuid=True), ForeignKey("tenants.id"), nullable=False)
    source_id = Column(UUID(as_uuid=True), ForeignKey("sources.id"), nullable=False)
    destination_id = Column(
        UUID(as_uuid=True), ForeignKey("destinations.id"), nullable=False
    )
    name = Column(String(255), nullable=False)
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    tenant = relationship("Tenant", back_populates="pipelines")
    source = relationship("Source", back_populates="pipelines")
    destination = relationship("Destination", back_populates="pipelines")
    tags = relationship("PipelineTag", back_populates="pipeline")
    app_configs = relationship("PipelineAppConfig", back_populates="pipeline")
    source_app_tables = relationship(
        "SourceAppTable",
        secondary="pipeline_source_app_tables",
        back_populates="pipelines",
    )


# Pipeline Tag
class PipelineTag(DeclarativeBase):
    __tablename__ = "pipeline_tags"

    pipeline_id = Column(
        UUID(as_uuid=True), ForeignKey("pipelines.id"), primary_key=True
    )  # Corrected from 'int' to UUID
    key = Column(String(255), primary_key=True)
    value = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    pipeline = relationship("Pipeline", back_populates="tags")


# Pipeline App Config
class PipelineAppConfig(DeclarativeBase):
    __tablename__ = "pipeline_app_configs"

    pipeline_id = Column(
        UUID(as_uuid=True), ForeignKey("pipelines.id"), primary_key=True
    )
    key = Column(String(255), primary_key=True)
    value = Column(String(255))
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    pipeline = relationship("Pipeline", back_populates="app_configs")


# Pipeline-SourceAppTable Association
class PipelineSourceAppTable(DeclarativeBase):
    __tablename__ = "pipeline_source_app_tables"

    pipeline_id = Column(
        UUID(as_uuid=True), ForeignKey("pipelines.id"), primary_key=True
    )
    table_id = Column(
        UUID(as_uuid=True), ForeignKey("source_app_tables.id"), primary_key=True
    )
    created_at = Column(DateTime, default=aware_utcnow())
    updated_at = Column(DateTime, default=aware_utcnow(), onupdate=aware_utcnow())

    pipeline = relationship("Pipeline", back_populates="source_app_table_associations")
    source_app_table = relationship(
        "SourceAppTable", back_populates="pipeline_associations"
    )
