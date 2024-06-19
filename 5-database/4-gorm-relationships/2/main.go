package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// relacionamento many to many: muitos produtos para muitas categorias

type (
	Category struct {
		ID       int `gorm:"primaryKey"`
		Name     string
		Products []*Product `gorm:"many2many:products_categories;"`
	}

	Product struct {
		ID         int `gorm:"primaryKey"`
		Name       string
		Price      float64
		Categories []*Category `gorm:"many2many:products_categories;"`
		gorm.Model
	}
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Category{}, &Product{})
	if err != nil {
		return
	}

	foodCategory := Category{Name: "Kitchen"}
	db.Create(&foodCategory)

	eletronicsCategory := Category{Name: "Eletronics"}
	db.Create(&eletronicsCategory)

	product := Product{
		Name:  "Frezzer",
		Price: 3000,
		Categories: []*Category{
			&foodCategory,
			&eletronicsCategory},
	}
	db.Create(&product)

	// consultando todos os produtos de uma categoria
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, p := range category.Products {
			fmt.Println("- ", p.Name, p.Price)
		}
	}

	// deletando para manter base limpa
	db.Migrator().DropTable(&Category{}, &Product{}, "products_categories")
}
