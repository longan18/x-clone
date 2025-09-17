package util

import "errors"

type AppError struct {
	RootErr error  `json:"-"`
	Message string `json:"message"`
	Log     string `json:"log"`
	Key     string `json:"key"`
}

func (e AppError) Error() string {
	return e.Message
}

func NewFullError(err error, msg, log, key string) AppError {
	return AppError{err, msg, log, key}
}

func NewError(err error, msg, key string) AppError {
	return AppError{err, msg, err.Error(), key}
}

func NewValidationError(err error) AppError {
	return AppError{
		err,
		"Invalid request, please check",
		err.Error(),
		"ErrValidaitonRequest",
	}
}

func NewCreateError(err error, key string) AppError {
	return AppError{
		err,
		"An error occurred during creation, please check again",
		err.Error(),
		key,
	}
}

func NewDuplicateError(msg, key string) AppError {
	err := errors.New(msg)

	return AppError{
		err,
		msg,
		err.Error(),
		key,
	}
}
