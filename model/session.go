package model

type Session struct {
	ID          uint                  `json:"id"`
	Date        string                `json:"phabricator"`
	UserID      uint                  `json:"user_id"`
	AirplaneID  uint                  `json:"airplane_id"`
	HasSolution []SessionHashSolution `json:"has_solution"`
	User        User
	Airplane    Airplane
}
