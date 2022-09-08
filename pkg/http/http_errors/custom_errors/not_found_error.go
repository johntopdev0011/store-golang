package customErrors

import (
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/http/http_errors/contracts"
	"github.com/pkg/errors"
	"net/http"
)

func NewNotFoundError(message string) error {
	ne := &notFoundError{
		WithStack: NewCustomErrorStack(nil, http.StatusNotFound, message),
	}

	return ne
}

func NewNotFoundErrorWrap(err error, message string) error {
	ne := &notFoundError{
		WithStack: NewCustomErrorStack(err, http.StatusNotFound, message),
	}

	return ne
}

type notFoundError struct {
	contracts.WithStack
}

type NotFoundError interface {
	contracts.WithStack
	IsNotFoundError() bool
	GetCustomError() CustomError
}

func (n *notFoundError) IsNotFoundError() bool {
	return true
}

func (n *notFoundError) GetCustomError() CustomError {
	return GetCustomError(n)
}

func IsNotFoundError(err error) bool {
	var notFoundError *notFoundError
	//us, ok := grpc_errors.Cause(err).(*notFoundError)
	if errors.As(err, &notFoundError) {
		return notFoundError.IsNotFoundError()
	}

	return false
}
