package util

type AppError struct {
	RootErr error  `json:"-"`
	Message string `json:"message"`
	Log     string `json:"log"`
	Key     string `json:"key"`
}

func NewUnprocessableEntity(err error) AppError {
	return AppError{
		err,
		"Invalid request, please check",
		err.Error(),
		"ErrValidaitonRequest",
	}
}
