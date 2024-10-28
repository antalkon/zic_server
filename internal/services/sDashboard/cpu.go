package sdashboard

import "github.com/gin-gonic/gin"

func LoadCPUPage(c *gin.Context) {
    c.HTML(200, "cpu_load.html", gin.H{
        "title":   "Загрузка процессора",
        "setting": true,
    })
}
