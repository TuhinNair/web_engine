package engine

type key int

const (
	// Original request URL
	ContextOriginalPath key = iota
	// Request start time
	ContextRequestStart
	// Any identifying datum
	ContextUserID
)
