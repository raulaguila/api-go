package repositories

import (
	"log"

	"github.com/raulaguila/api-go/api/domain"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Insert(profile *domain.Profile) (*domain.Profile, error)
}

type ProfileRepositoryDB struct {
	DB *gorm.DB
}

func (repo *ProfileRepositoryDB) Insert(profile *domain.Profile) (*domain.Profile, error) {
	if err := repo.DB.Create(&profile).Error; err != nil {
		log.Fatalf("Error to persist profile: %v", err)
		return profile, err
	}

	return profile, nil
}
