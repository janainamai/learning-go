package database

import "github.com/janainamai/study-api-go/internal/entity"

type (
	UserDatabaseInterface interface {
		Create(user *entity.User) error
		FindByEmail(email string) (*entity.User, error)
	}

	ProductDatabaseInterface interface {
		Create(product *entity.Product) error
		FindAll(page, limit int, sort string) ([]*entity.Product, error)
		FindByID(id string) (*entity.Product, error)
		Update(product *entity.Product) error
		Delete(id string) error
	}
)
