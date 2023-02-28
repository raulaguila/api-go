package main

import (
	"fmt"
	"log"

	"github.com/raulaguila/api-go/api/application/repositories"
	"github.com/raulaguila/api-go/api/domain"
	"github.com/raulaguila/api-go/api/framework/utils"
)

func main() {

	user := domain.User{
		Name:     "João de Barro",
		Email:    "joao@barro.com",
		Password: "12345678",
	}

	userRepo := repositories.UserRepositoryDB{DB: utils.ConnectDB()}
	if result, err := userRepo.Insert(&user); err != nil {
		log.Fatalf("Error to create user: %v", err)
	} else {
		fmt.Println("User created succesfully!", result)
	}
}
