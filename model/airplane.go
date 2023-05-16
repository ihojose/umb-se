package model

type Airplane struct {
	ID          uint      `json:"id"`
	Phabricator string    `json:"phabricator"`
	ModelName   string    `json:"model_name"`
	ModelLine   string    `json:"model_line"`
	ModelUpdate string    `json:"model_update"`
	Sessions    []Session `json:"sessions"`
}
