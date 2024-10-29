package sdata

import "github.com/gin-gonic/gin"

func LoadDataPage(c *gin.Context) {
	c.HTML(200, "data.html", gin.H{
		"title":   "Данные",
		"setting": false,
	})
}
