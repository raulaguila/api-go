package domain

import (
	"errors"
	"log"
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Name            *string `json:"name"`
	Email           *string `json:"email"`
	Password        *string `json:"password"`
	PasswordConfirm *string `json:"password_confirm"`
	ProfileId       *uint   `json:"profile_id"`
}

type User struct {
	Base
	Name      string   `json:"name" gorm:"column:name;type:varchar(255)"`
	Email     string   `json:"email" gorm:"column:email;type:varchar(255);unique;index"`
	Password  string   `json:"-" gorm:"column:password;type:varchar(255)"`
	Token     string   `json:"-" gorm:"column:token;type:varchar(255);unique;index"`
	ProfileId uint     `json:"-" gorm:"column:profile_id;"`
	Profile   *Profile `json:"profile" gorm:"foreignKey:ProfileId"`
}

func (u *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Prepare() error {
	if !strings.HasPrefix(u.Password, "$") || len(u.Password) != 60 {
		if hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
			log.Fatalf("Error during the password generation: %v\n", err)
			return err
		} else {
			u.Password = string(hashed)
		}
	}

	if len(u.Token) == 0 {
		u.Token = uuid.NewV4().String()
	}

	return u.validate()
}

func (u *User) validate() error {
	return nil
}

func (u *User) From(data CreateUser) error {
	if data.Password != nil && data.PasswordConfirm != nil {
		if *data.Password != *data.PasswordConfirm {
			return errors.New("passwords do not match")
		}
		u.Password = *data.Password
	}

	if data.Email != nil {
		u.Email = *data.Email
	}

	if data.Name != nil {
		u.Name = *data.Name
	}

	if data.ProfileId != nil {
		u.ProfileId = *data.ProfileId
	}

	return nil
}
