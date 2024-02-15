package model

type Response struct {
	Status  int `json:"status"`
	Message any `json:"message"`
}
