package customerror

import "fmt"

// RecordNotFound is a custom error which could be thrown if API or DB records were not found
type RecordNotFound struct {
	Type             string
	IdentifierColumn string
	Identifier       string
	Err              error
	PrintOriginalErr bool
}

// Error prints the full error message as string
func (e *RecordNotFound) Error() string {
	errMsg := fmt.Sprintf("Could not found %s identified with %s %s", e.Type, e.IdentifierColumn, e.Identifier)

	if e.PrintOriginalErr {
		errMsg = fmt.Sprintf("Could not found %s identified with %s %s\n(%s)", e.Type, e.IdentifierColumn, e.Identifier, e.Err.Error())
	}

	return errMsg
}
