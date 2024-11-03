package handler

import (
	scloud "github.com/antalkon/zic_server/internal/services/sCloud"
	"github.com/gin-gonic/gin"
)

func CloudPage(c *gin.Context) {
	c.HTML(200, "cloud.html", gin.H{
		"title":   "Облако",
		"page":    "cloud",
		"setting": true,
	})
}

func CClassPage(c *gin.Context) {
	scloud.ClassPage(c)
}

func CreateClass(c *gin.Context) {
	scloud.CreateClass(c)
}

func DelClass(c *gin.Context) {
	scloud.DelClass(c)
}

func CTaskPage(c *gin.Context) {
	scloud.TaskPage(c)
}

func GetClasses(c *gin.Context) {
	scloud.GetClasses(c)
}

func NewTask(c *gin.Context) {
	scloud.NewTask(c)
}

func AllTasks(c *gin.Context) {
	scloud.AllTasks(c)
}

func DelTask(c *gin.Context) {
	scloud.DelTask(c)
}

func GetTaskClass(c *gin.Context) {
	scloud.TaskClass(c)
}

func TaskDelivery(c *gin.Context) {
	scloud.TaskDelivery(c)
}

func DeliveryAll(c *gin.Context) {
	scloud.DeliveryAll(c)
}

func DeliveryChecked(c *gin.Context) {
	scloud.Checked(c)
}

func DelTaskDel(c *gin.Context) {
	scloud.DelTaskDel(c)
}
