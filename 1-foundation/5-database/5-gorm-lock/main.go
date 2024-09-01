package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// representando lock pessimista

type (
	Category struct {
		ID   int `gorm:"primaryKey"`
		Name string
	}
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Category{})
	if err != nil {
		return
	}

	tx := db.Begin()

	eletronicsCategory := Category{Name: "Eletronics"}
	db.Create(&eletronicsCategory)

	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}

	c.Name = "Books"
	tx.Debug().Save(&c)
	tx.Commit()

	// deletando para manter base limpa
	db.Migrator().DropTable(&Category{})
}
