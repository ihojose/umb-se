package model

type Session struct {
	ID          uint                  `json:"id"`
	Date        string                `json:"date"`
	UserID      uint                  `json:"user_id"`
	AirplaneID  uint                  `json:"airplane_id"`
	HasSolution []SessionHashSolution `json:"has_solution"`
	History     []History             `json:"history"`
}
