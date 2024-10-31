package serrors

import "github.com/gin-gonic/gin"

func Page403(c *gin.Context) {
	c.HTML(403, "403.html", gin.H{
		"title": "Ошибка доступа",
	})
}
