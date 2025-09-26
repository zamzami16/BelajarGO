package exception

import (
	"github.com/sirupsen/logrus"
)

// ContextualError wraps an error with logger context
type ContextualError struct {
	Err    error
	Logger *logrus.Entry
}

func (ce ContextualError) Error() string {
	return ce.Err.Error()
}

// NewContextualError creates a new contextual error
func NewContextualError(err error, logger *logrus.Entry) ContextualError {
	return ContextualError{
		Err:    err,
		Logger: logger,
	}
}

// ContextualNotFoundError with logger context
type ContextualNotFoundError struct {
	NotFoundError
	Logger *logrus.Entry
}

func NewContextualNotFoundError(message string, logger *logrus.Entry) ContextualNotFoundError {
	return ContextualNotFoundError{
		NotFoundError: NotFoundError{Error: message},
		Logger:        logger,
	}
}
