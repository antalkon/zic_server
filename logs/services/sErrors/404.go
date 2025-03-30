package serrors

import "github.com/gin-gonic/gin"

func Page404(c *gin.Context) {
	c.HTML(404, "404.html", gin.H{
		"title": "Ошибка доступа",
	})
}
