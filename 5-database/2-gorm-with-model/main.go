package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}

	// INSERINDO UM REGISTRO
	db.Create(&Product{
		Name:  "Macbook",
		Price: 10000.00,
	})
	fmt.Println("1 Produto inserido com sucesso")

	// CONSULTANDO UM PRODUTO POR ID
	var productOne Product
	db.First(&productOne)
	fmt.Printf("Produto encontrado por identificador, ID: %d, Name: %s, Price: %.2f\n", productOne.ID, productOne.Name, productOne.Price)

	// ALTERANDO REGISTRO POR ID
	productOne.Name = "New name"
	productOne.Price = 99.99
	db.Save(productOne)
	fmt.Println("Produto alterado com sucesso")

	var productUpdated Product
	db.First(&productUpdated, productOne.ID)
	fmt.Printf("Consultando produto que foi alterado ID: %d, Name: %s, Price: %.2f\n",
		productUpdated.ID,
		productUpdated.Name,
		productUpdated.Price)

	// REMOVENDO O REGISTRO
	db.Delete(&productUpdated)
	fmt.Println("Produto removido com sucesso")

}
