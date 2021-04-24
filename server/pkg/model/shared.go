package model

// Status is the status response
type Status struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// Error is the standard error response
type Error struct {
	Message string `json:"message"`
}
