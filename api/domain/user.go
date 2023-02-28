package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255);unique;index"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Token    string `json:"token" gorm:"type:varchar(255);unique;index"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) Prepare() error {
	if hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		log.Fatalf("Error during the password generation: %v\n", err)
		return err
	} else {
		u.Password = string(hashed)
	}
	u.Token = uuid.NewV4().String()

	return u.validate()
}

func (u *User) validate() error {
	return nil
}
