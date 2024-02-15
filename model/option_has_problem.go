package model

type OptionHasProblem struct {
	OptionID  uint    `json:"option_id"`
	ProblemID string  `json:"problem_id"`
	Problem   Problem `json:"problem" gorm:"foreignKey:ProblemID;references:ID"`
}
