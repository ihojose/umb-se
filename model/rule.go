package model

type Rule struct {
	ID      string   `json:"id" gorm:"primaryKey"`
	Rule    string   `json:"rule"`
	Options []Option `json:"options" gorm:"foreignKey:RuleID;references:ID"`
}
