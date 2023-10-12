package ini

import (
	"fmt"
)

// ErrDelimiterNotFound indicates the error type of no delimiter is found which there should be one.
type ErrDelimiterNotFound struct {
	Line string
}

// IsErrDelimiterNotFound returns true if the given error is an instance of ErrDelimiterNotFound.
func IsErrDelimiterNotFound(err error) bool {
	_, ok := err.(ErrDelimiterNotFound)
	return ok
}

func (err ErrDelimiterNotFound) Error() string {
	return fmt.Sprintf("key-value delimiter not found: %s", err.Line)
}

// ErrEmptyKeyName indicates the error type of no key name is found which there should be one.
type ErrEmptyKeyName struct {
	Line string
}

// IsErrEmptyKeyName returns true if the given error is an instance of ErrEmptyKeyName.
func IsErrEmptyKeyName(err error) bool {
	_, ok := err.(ErrEmptyKeyName)
	return ok
}

func (err ErrEmptyKeyName) Error() string {
	return fmt.Sprintf("empty key name: %s", err.Line)
}
