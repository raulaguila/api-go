package domain

type Profile struct {
	Base
	Name string `json:"name" gorm:"type:varchar(255);unique"`
}
