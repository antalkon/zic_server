package handler

import "github.com/gin-gonic/gin"

func TasksPage(c *gin.Context) {
	c.HTML(200, "tasks_u.html", gin.H{
		"title": "Задачи",
	})
}

func TasksDPage(c *gin.Context) {
	c.HTML(200, "tasksD_u.html", gin.H{
		"title": "Загрузка заданий",
	})
}
