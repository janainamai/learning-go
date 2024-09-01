package database

import (
	"github.com/janainamai/study-api-go/internal/entity"
	"gorm.io/gorm"
)

type ProductDatabase struct {
	db *gorm.DB
}

func NewProductDatabase(db *gorm.DB) *ProductDatabase {
	return &ProductDatabase{db: db}
}

func (p *ProductDatabase) Create(product *entity.Product) error {
	return p.db.Create(product).Error
}

func (p *ProductDatabase) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		var products []*entity.Product
		err := p.db.Limit(limit).Offset((page - 1) * limit).Order("create_at " + sort).Find(&products).Error
		return products, err
	}

	var products []*entity.Product
	err := p.db.Limit(limit).Order("create_at " + sort).Find(&products).Error
	return products, err
}

func (p *ProductDatabase) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.db.First(&product, "id = ?", id).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}

		return nil, err
	}

	return &product, nil
}

func (p *ProductDatabase) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}

	return p.db.Save(product).Error
}

func (p *ProductDatabase) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}

	return p.db.Delete(product).Error
}
