package sauth

import "github.com/gin-gonic/gin"

func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "Вход",
	})
}
