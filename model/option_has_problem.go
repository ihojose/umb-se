package model

type OptionHasProblem struct {
	OptionID  uint   `json:"option_id"`
	ProblemID string `json:"problem_id"`
	Option    Option
	Problem   Problem
}
