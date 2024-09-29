package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User структура представляет пользователя
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Пример базы данных пользователей
var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Smith"},
}

// GetUsers возвращает список пользователей
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// CreateUser создает нового пользователя
func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}
