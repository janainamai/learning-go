package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// relacionamento one to many: um produto para muitas categorias
// relacionamento one to one: um serial number para um produto
// relacionamento many to one: uma categoria para muitos produtos

type (
	Category struct {
		ID       int `gorm:"primaryKey"`
		Name     string
		Products []*Product
	}

	SerialNumber struct {
		ID        int `gorm:"primaryKey"`
		Number    string
		ProductID int // associação has one - one to one - one serial number to one product
	}

	Product struct {
		ID           int `gorm:"primaryKey"`
		Name         string
		Price        float64
		CategoryID   int           // chave estrangeira
		Category     *Category     // associação belongs to - many to one - many products to one category
		SerialNumber *SerialNumber // associação has one - one to one - one serial number to one product
		gorm.Model
	}
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})
	if err != nil {
		return
	}

	// create category
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	// create product
	product := Product{Name: "Macbook", Price: 10000.00, CategoryID: category.ID}
	db.Create(&product)

	// create serial number
	serialNumber := SerialNumber{Number: "123456", ProductID: product.ID}
	db.Create(&serialNumber)

	// consultando todos os produtos
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

	// consultando todos os produtos de uma categoria
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, p := range category.Products {
			fmt.Println("- ", p.Name, p.Price, p.SerialNumber.Number)
		}
	}

	// deletando para manter base limpa
	db.Migrator().DropTable(&Category{}, &Product{}, &SerialNumber{})

}
