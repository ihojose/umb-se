package model

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	IdType    string `json:"id_type"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Password  string `json:"password" gorm:"-"`
	HashPass  string
	HashKey   string
	HashToken string
	Role      int32     `json:"role"`
	Sessions  []Session `json:"sessions"`
}
