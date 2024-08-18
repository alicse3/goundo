package util

import "fmt"

// AppErr is a custom error type that wraps an error message and the underlying error
type AppErr struct {
	Message string
	Err     error
}

// Error implements the error interface for AppErr
// It returns a formatted string containing the error message and the underlying error
func (ae *AppErr) Error() string {
	if ae.Err != nil {
		return fmt.Sprintf("Error => Message: %s, Err: %v\n", ae.Message, ae.Err)
	}
	return fmt.Sprintf("Error => Message: %s\n", ae.Message)
}
