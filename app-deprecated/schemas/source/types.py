import enum


class SQLSourceBinaryHandlingMode(str, enum.Enum):
    BYTES = "bytes"
    BASE64 = "base64"
    BASE64_URL_SAFE = "base64-url-safe"
    HEX = "hex"
