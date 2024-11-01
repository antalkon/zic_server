package handler

import (
	sauth "github.com/antalkon/zic_server/internal/services/sAuth"
	"github.com/gin-gonic/gin"
)

// post запрос на вход
func Login(c *gin.Context) {
	sauth.Login(c)
}

// страница входа
func LoginPage(c *gin.Context) {
	sauth.LoginPage(c)
}
