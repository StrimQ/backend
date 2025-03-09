import structlog
from fastapi import FastAPI, Request, Response
from fastapi.middleware.cors import CORSMiddleware
from fastapi.openapi.docs import get_swagger_ui_html

from .logging import configure_logging
from .settings import settings

app = FastAPI(
    title="Real-time data streaming with FastAPI and Temporal",
    version="0.1.0",
    debug=settings.ENVIRONMENT == "development",
)
configure_logging()
LOG: structlog.stdlib.BoundLogger = structlog.get_logger()


@app.middleware("http")
async def log_requests(request: Request, call_next):
    LOG.info("Request", method=request.method, url=str(request.url))
    response: Response = await call_next(request)
    LOG.info("Response", status_code=response.status_code)
    return response


# CORS for frontend
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Replace with specific origins in production
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# app.include_router(hello_router)


# https://github.com/tiangolo/fastapi/issues/4211#issuecomment-983600161
@app.get("/docs", include_in_schema=False)
async def custom_swagger_ui_html(request: Request):
    LOG.info(f"Swagger UI URL: {app.openapi_url}")
    return get_swagger_ui_html(
        openapi_url=str(app.openapi_url),
        title="API",
    )
