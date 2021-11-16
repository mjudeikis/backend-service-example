package models

type APIError string

const (
	ErrEmailRequired  APIError = "Email is required"
	ErrPasswordFormat APIError = "Your password is weak!"
)

func (s APIError) Error() string {
	return string(s)
}

// CloudError represents a cloud error.
type CloudError struct {
	// The status code.
	StatusCode int `json:"-"`

	// An error response from the service.
	CloudErrorBody string `json:"error,omitempty"`
}
