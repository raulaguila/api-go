package repositories

import (
	"log"

	"github.com/raulaguila/api-go/api/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDB struct {
	DB *gorm.DB
}

func (repo *UserRepositoryDB) Insert(user *domain.User) (*domain.User, error) {
	if err := user.Prepare(); err != nil {
		log.Fatalf("Error during the user validation: %v", err)
		return user, err
	}

	if err := repo.DB.Create(&user).Error; err != nil {
		log.Fatalf("Error to persist user: %v", err)
		return user, err
	}

	return user, nil
}
