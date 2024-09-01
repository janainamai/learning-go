package database

import (
	"testing"

	"github.com/janainamai/study-api-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("Jana", "jana@hotmail.com", "12345")
	if err != nil {
		t.Error(err)
	}

	userDB := NewUserDatabase(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	var userCreated entity.User
	err = db.First(&userCreated, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userCreated.ID)
	assert.Equal(t, user.Name, userCreated.Name)
	assert.Equal(t, user.Email, userCreated.Email)
	assert.Equal(t, user.Password, userCreated.Password)

	var allData []entity.User
	db.Find(&allData)
	db.Delete(allData)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("Jana", "jana@hotmail.com", "12345")
	if err != nil {
		t.Error(err)
	}
	print(user.Password)

	userDB := NewUserDatabase(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)

	var allData []entity.User
	db.Find(&allData)
	db.Delete(allData)
}
