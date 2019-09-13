package app

// ErrorResp represents error model, the default error model for all responses.
type ErrorResp struct {
	Name    string        `json:"name,omitempty"`
	Message string        `json:"message,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
}

// ErrorDetail represents error detail model.
type ErrorDetail struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}
