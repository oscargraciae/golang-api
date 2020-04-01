package repository

import (
	"api-app/api/models"
	"api-app/api/security"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) SaveUser(user models.User) (models.User, error) {
	err := u.DB.Debug().Create(&user).Error
	return user, err
}

func (u *UserRepository) GetAll() []models.User {
	users := []models.User{}
	u.DB.Debug().Find(&users)
	return users
}

func (u* UserRepository) GetUserLogin(email, password string) (models.User, error) {
	user := models.User{}
	u.DB.Debug().First(&user, "email = ?", email)

	err := security.VerifyPassword(user.Password, password)
	fmt.Println("ERROS---->", err)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
