package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/janainamai/study-api-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDatabase(db)

	product, err := entity.NewProduct("Product 1", 10.0)
	if err != nil {
		t.Error(err)
	}

	err = productDB.Create(product)
	assert.Nil(t, err)

	var productSaved entity.Product
	err = db.First(&productSaved, "id = ?", product.ID).Error
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, productSaved)
	assert.Equal(t, product.ID, productSaved.ID)
	assert.Equal(t, product.Name, productSaved.Name)
	assert.Equal(t, product.Price, productSaved.Price)

	var allData []entity.Product
	db.Find(&allData)
	db.Delete(allData)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDatabase(db)

	for i := 1; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*50)
		if err != nil {
			t.Error(err)
		}

		err = db.Create(product).Error
		assert.Nil(t, err)
	}

	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 24", products[3].Name)

	var allData []entity.Product
	db.Find(&allData)
	db.Delete(allData)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDatabase(db)

	product, err := entity.NewProduct("Product 1", 10.0)
	if err != nil {
		t.Error(err)
	}

	err = productDB.Create(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

	var allData []entity.Product
	db.Find(&allData)
	db.Delete(allData)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDatabase(db)

	product, err := entity.NewProduct("Product 1", 10.0)
	if err != nil {
		t.Error(err)
	}

	err = productDB.Create(product)
	assert.Nil(t, err)

	product.Name = "Product updated"
	product.Price = 500.50
	err = productDB.Update(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

	var allData []entity.Product
	db.Find(&allData)
	db.Delete(allData)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProductDatabase(db)

	product, err := entity.NewProduct("Product 1", 10)
	if err != nil {
		t.Error(err)
	}

	err = productDB.Create(product)
	assert.Nil(t, err)

	// err = productDB.Delete(product.ID.String())
	// assert.Nil(t, err)

	// product, err = productDB.FindByID(product.ID.String())
	// assert.NotNil(t, err)
	// assert.Nil(t, product)

	var allData []entity.Product
	db.Find(&allData)
	db.Delete(allData)
}
