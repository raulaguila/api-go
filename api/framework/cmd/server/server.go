package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/raulaguila/api-go/api/application/repositories"
	"github.com/raulaguila/api-go/api/domain"
	"github.com/raulaguila/api-go/api/framework/utils"
)

func main() {

	db := utils.ConnectDB()
	profile := domain.Profile{
		Name: "admin",
	}
	profileRepo := repositories.ProfileRepositoryDB{DB: db}
	if result, err := profileRepo.Insert(&profile); err != nil {
		fmt.Printf("Error to create profile: %v", err)
	} else {
		fmt.Println("Profile created succesfully!", result)
	}

	name := "João do Barro"
	email := "joao@barro.com"
	password := "12345678"
	password2 := "12345678"
	user := domain.CreateUser{
		Name:            &name,
		Email:           &email,
		Password:        &password,
		PasswordConfirm: &password2,
		ProfileId:       &profile.Id,
	}

	var newUser domain.User
	if err := newUser.From(user); err != nil {
		log.Fatalf("Error to convert user: %v", err)
	}

	userRepo := repositories.UserRepositoryDB{DB: db}
	if err := userRepo.Insert(&newUser); err != nil {
		log.Fatalf("Error to create user: %v", err)
	} else {
		fmt.Println("User created succesfully!", newUser.Profile.Name)
	}

	js, _ := json.Marshal(newUser)
	fmt.Println(string(js))

	var users []domain.User
	userRepo.RequestAll(&users)
	js, _ = json.Marshal(users)
	fmt.Println(string(js))
}
