package controllers

import (
	"api-app/api/auth"
	"api-app/api/database"
	"api-app/api/models"
	"api-app/api/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	r := repository.UserRepository{DB: database.Connect()}
	users := r.GetAll()
	c.JSON(200, gin.H{
		"users": users,
	})
}

func CreateUser(c *gin.Context) {
	r := repository.UserRepository{DB: database.Connect()}
	var user models.User
	c.BindJSON(&user)
	user, err := r.SaveUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Usuario creado",
	})
}

func Login(c *gin.Context) {
	user := models.User{}
	r := repository.UserRepository{DB: database.Connect()}

	c.BindJSON(&user)
	user, err := r.GetUserLogin(user.Email, user.Password)
	fmt.Println("ERror----->", err)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := auth.GenToken(user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, token)
}