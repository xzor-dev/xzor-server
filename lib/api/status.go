package api

const (
	// StatusSuccess indicates that the request was successful.
	StatusSuccess Status = "success"

	// StatusFail indicates that the request failed.
	StatusFail Status = "fail"

	// StatusError indicates that the request generated an error.
	StatusError Status = "error"
)

// Status is a string indicating a request's status.
type Status string
