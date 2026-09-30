package cdksopssecrets


type ResourceType string

const (
	ResourceType_SECRET ResourceType = "SECRET"
	ResourceType_SECRET_RAW ResourceType = "SECRET_RAW"
	ResourceType_SECRET_BINARY ResourceType = "SECRET_BINARY"
	ResourceType_PARAMETER ResourceType = "PARAMETER"
	ResourceType_PARAMETER_MULTI ResourceType = "PARAMETER_MULTI"
)

