package model

type Problem struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	UrgencyLevel int32    `json:"urgency_level"`
	OptionId     uint     `json:"option_id" gorm:"-"`
	Solution     Solution `json:"solution" gorm:"foreignKey:ProblemID;references:ID"`
}
