package repositories

import (
	"fmt"

	"github.com/raulaguila/api-go/api/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	RequestAll(user *[]domain.User) error
	Request(user *domain.User, id uint) error
	Insert(user *domain.User) error
	Delete(id uint) error
}

type UserRepositoryDB struct {
	DB *gorm.DB
}

func (repo *UserRepositoryDB) RequestAll(user *[]domain.User) error {
	if err := repo.DB.Preload("Profile").Find(&user).Error; err != nil {
		fmt.Printf("Error during the user request: %v", err)
		return err
	}

	return nil
}

func (repo *UserRepositoryDB) Request(user *domain.User, id uint) error {
	if err := repo.DB.Preload("Profile").First(&user, id).Error; err != nil {
		fmt.Printf("Error during the user request: %v", err)
		return err
	}

	return nil
}

func (repo *UserRepositoryDB) Insert(user *domain.User) error {
	if err := user.Prepare(); err != nil {
		fmt.Printf("Error during the user validation: %v", err)
		return err
	}

	if err := repo.DB.Create(&user).Error; err != nil {
		fmt.Printf("Error to persist user: %v", err)
		return err
	}

	return repo.Request(user, user.Id)
}

func (repo *UserRepositoryDB) Delete(id uint) error {
	if err := repo.DB.Delete(&domain.User{}, id).Error; err != nil {
		fmt.Printf("Error to delete user: %v", err)
		return err
	}

	return nil
}
