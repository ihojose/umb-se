package model

type History struct {
	ID        uint   `json:"id"`
	Datetime  string `json:"datetime"`
	SessionID uint   `json:"session_id"`
	OptionID  uint   `json:"option_id"`
}
