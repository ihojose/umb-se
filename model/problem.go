package model

type Problem struct {
	ID           string             `json:"id"`
	Description  string             `json:"description"`
	UrgencyLevel int32              `json:"urgency_level"`
	HasOption    []OptionHasProblem `json:"has_option"`
}
