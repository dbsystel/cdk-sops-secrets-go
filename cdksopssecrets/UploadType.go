package cdksopssecrets


type UploadType string

const (
	// Pass the secret data inline (base64 encoded and compressed).
	UploadType_INLINE UploadType = "INLINE"
	// Upload the secret data as asset.
	UploadType_ASSET UploadType = "ASSET"
)

