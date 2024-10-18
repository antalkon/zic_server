package handler

import (
	sauth "github.com/antalkon/zic_server/internal/services/sAuth"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	sauth.Login(c)
}

func LoginPage(c *gin.Context) {
	sauth.LoginPage(c)
}
