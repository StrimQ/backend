package schemas

// SQLSourceBinaryHandlingMode represents the binary handling mode enum.
type SQLSourceBinaryHandlingMode string

const (
	Bytes         SQLSourceBinaryHandlingMode = "bytes"
	Base64        SQLSourceBinaryHandlingMode = "base64"
	Base64URLSafe SQLSourceBinaryHandlingMode = "base64-url-safe"
	Hex           SQLSourceBinaryHandlingMode = "hex"
)
