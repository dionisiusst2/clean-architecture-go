package errors

import (
	"fmt"
	"strings"
)

type HttpError interface {
	WithOperation(string) HttpError
	PrintFailedOperations()
	GetStatusCode() int
	Error() string
}

type httpError struct {
	operations   []string
	StatusCode   int
	ErrorMessage string
}

func NewHttpError(operation string, httpStatusCode int, errorMessage string) HttpError {
	return &httpError{[]string{operation}, httpStatusCode, errorMessage}
}

func (e *httpError) WithOperation(operation string) HttpError {
	e.operations = append(e.operations, operation)
	return e
}

func (e *httpError) PrintFailedOperations() {
	fmt.Println("\nERROR OPERATIONS: ", strings.Join(e.operations[:], " - "))
}

func (e *httpError) GetStatusCode() int {
	return e.StatusCode
}

func (e *httpError) Error() string {
	return e.ErrorMessage
}
