package router

import "fmt"

// ErrNilProject is returned when a nil project is passed to Route.
var ErrNilProject = fmt.Errorf("verso: project cannot be nil")

// ErrWorkflowNotFound is returned when a requested workflow is not found.
type ErrWorkflowNotFound struct {
	Name string
}

func (e *ErrWorkflowNotFound) Error() string {
	return fmt.Sprintf("verso: workflow %q not found", e.Name)
}

// ErrInvalidStrategy is returned when an unknown strategy is specified.
type ErrInvalidStrategy struct {
	Strategy string
}

func (e *ErrInvalidStrategy) Error() string {
	return fmt.Sprintf("verso: invalid routing strategy %q", e.Strategy)
}