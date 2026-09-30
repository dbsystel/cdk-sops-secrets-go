package cdksopssecrets


type RawOutput string

const (
	// Parse the secret as a string.
	RawOutput_STRING RawOutput = "STRING"
	// Parse the secret as a binary.
	RawOutput_BINARY RawOutput = "BINARY"
)

