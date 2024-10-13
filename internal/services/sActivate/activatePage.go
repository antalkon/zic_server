package sactivate

import "github.com/gin-gonic/gin"

func ActPage(c *gin.Context) {
	c.HTML(200, "activate.html", gin.H{
		"title": "Вход",
	})
}
