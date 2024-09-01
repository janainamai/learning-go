package database

import (
	"github.com/janainamai/study-api-go/internal/entity"
	"gorm.io/gorm"
)

type UserDatabase struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) *UserDatabase {
	return &UserDatabase{db: db}
}

func (u *UserDatabase) Create(user *entity.User) error {
	return u.db.Create(user).Error
}

func (u *UserDatabase) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
