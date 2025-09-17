package util

type AppSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewCreateSuccess(data interface{}) AppSuccess {
	return AppSuccess{
		"Created successfully",
		data,
	}
}

func NewGetSuccess(data interface{}) AppSuccess {
	return AppSuccess{
		"Get data successfully",
		data,
	}
}
