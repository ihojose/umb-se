package model

type Option struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Description string         `json:"description"`
	RuleID      string         `json:"rule_id"`
	NextRule    OptionNextRule `json:"next_rule" gorm:"foreignKey:OptionID;references:ID"`
}
