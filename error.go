package valid

import "fmt"

// ValidationErrors is errors occurred when validation.
type ValidationErrors struct {
	Errors []error
}

// Error returns string representation for this error.
func (err *ValidationErrors) Error() string {
	return fmt.Sprintf("validation errors: %v", err.Errors)
}

// First returns first validation error.
func (err *ValidationErrors) First() error {
	return err.Errors[0]
}
