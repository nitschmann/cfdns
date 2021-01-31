package customerror

import "fmt"

// RecordNotFoound is a custom error which could be thrown if API or DB records were not found
type RecordNotFound struct {
	Type             string
	IdentifierColumn string
	Identifier       string
	Err              error
}

// Error prints the full error message as string
func (e *RecordNotFound) Error() string {
	return fmt.Sprintf("Could not found %s identified with %s %s", e.Type, e.IdentifierColumn, e.Identifier)
}
