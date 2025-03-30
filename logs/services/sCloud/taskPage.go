package scloud

import "github.com/gin-gonic/gin"

func TaskPage(c *gin.Context) {
	c.HTML(200, "cl_tasks.html", gin.H{
		"title": "Задачи",
	})
}
