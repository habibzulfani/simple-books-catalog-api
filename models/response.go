package models

type Response struct {
	Status  string `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Data    any    `json:"data,omitempty" form:"data"`
}
