package valid

import "fmt"

type ValidationErrors struct {
	Errors []error
}

func (err *ValidationErrors) Error() string {
	return fmt.Sprintf("validation errors: %v", err.Errors)
}

func (err *ValidationErrors) First() error {
	return err.Errors[0]
}
