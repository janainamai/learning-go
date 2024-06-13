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
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
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

	// INSERINDO MAIS REGISTROS
	products := []Product{
		{Name: "Mouse", Price: 200.00},
		{Name: "Teclado", Price: 300.00},
		{Name: "Monitor", Price: 1000.00},
		{Name: "Book", Price: 40.50},
	}
	db.Create(&products)
	fmt.Println("3 Produto inserido com sucesso")

	// CONSULTANDO UM PRODUTO POR ID
	var productOne Product
	db.First(&productOne, 1)
	fmt.Printf("Produto encontrado por identificador, ID: %d, Name: %s, Price: %.2f\n", productOne.ID, productOne.Name, productOne.Price)

	// CONSULTANDO UM PRODUTO PELO NOME
	var productMouse Product
	db.First(&productMouse, "name = ?", "Mouse")
	fmt.Printf("Produto encontrado por nome, ID: %d, Name: %s, Price: %.2f\n", productMouse.ID, productMouse.Name, productMouse.Price)

	// CONSULTANDO TODOS OS PRODUTOS
	var allProducts []Product
	db.Find(&allProducts)
	fmt.Println("\nTodos os produtos foram buscados:")
	for _, p := range allProducts {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

	// CONSULTANDO TODOS OS PRODUTOS COM PAGINAÇÃO
	var allProductsPageable []Product
	db.Limit(10).Offset(3).Find(&allProductsPageable)
	fmt.Println("\nTodos os produtos foram buscados com paginação:")
	for _, p := range allProductsPageable {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

	// CONSULTANDO DADOS COM WHERE EXPLÍCITO
	var productsWhere []Product
	db.Where("price > 800").Find(&productsWhere)
	fmt.Println("\nTodos os produtos com preço maior que 800 foram buscados:")
	for _, p := range productsWhere {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

	var productsLike []Product
	db.Where("name LIKE ?", "%book%").Find(&productsLike)
	fmt.Println("\nTodos os produtos com nome contendo 'book' foram buscados:")
	for _, p := range productsLike {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

	// ALTERANDO REGISTRO POR ID
	id := 1
	var productToUpdate Product
	db.First(&productToUpdate, id)
	fmt.Printf("Produto a ser alterado ID: %d, Name: %s, Price: %.2f\n",
		productToUpdate.ID,
		productToUpdate.Name,
		productToUpdate.Price)

	productToUpdate.Name = "New name"
	productToUpdate.Price = 99.99
	db.Save(productToUpdate)
	fmt.Println("Produto alterado com sucesso")

	var productUpdated Product
	db.First(&productUpdated, id)
	fmt.Printf("Consultando produto que foi alterado ID: %d, Name: %s, Price: %.2f\n",
		productToUpdate.ID,
		productToUpdate.Name,
		productToUpdate.Price)

	// REMOVENDO O REGISTRO POR ID
	db.Delete(productUpdated)
	fmt.Println("Produto removido com sucesso")

	// REMOVENDO TODOS OS REGISTROS
	var allProductsToDelete []Product
	db.Find(&allProductsToDelete)
	db.Delete(allProductsToDelete)
	fmt.Println("Todos os produtos foram deletados com sucesso")
}
