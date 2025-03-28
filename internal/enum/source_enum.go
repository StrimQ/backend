package enum

// SourceBinaryHandlingMode represents the binary handling mode enum.
type SourceBinaryHandlingMode string

const (
	SourceBinaryHandlingMode_Bytes         SourceBinaryHandlingMode = "bytes"
	SourceBinaryHandlingMode_Base64        SourceBinaryHandlingMode = "base64"
	SourceBinaryHandlingMode_Base64urlsafe SourceBinaryHandlingMode = "base64-url-safe"
	SourceBinaryHandlingMode_Hex           SourceBinaryHandlingMode = "hex"
)

// SourceSSLMode represents the SSL mode enum.
type SourceSSLMode string

const (
	SourceSSLMode_Disable SourceSSLMode = "disable"
	SourceSSLMode_Require SourceSSLMode = "require"
)

// SourceEngine represents the source engine enum.
type SourceEngine string

const (
	SourceEngine_Mysql      SourceEngine = "mysql"
	SourceEngine_Postgresql SourceEngine = "postgresql"
)
