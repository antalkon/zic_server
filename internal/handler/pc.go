package handler

import (
	spc "github.com/antalkon/zic_server/internal/services/sPc"
	"github.com/gin-gonic/gin"
)

// post создания нового пк
func AddNewPc(c *gin.Context) {
	spc.AddPc(c)
}

// post количества пк
func PcCount(c *gin.Context) {
	spc.PcCount(c)
}

// пинг сервера от пк
func ServerPing(c *gin.Context) {
	spc.PcPing(c)
}

// post выключения пк
func PcOff(c *gin.Context) {
	spc.OffPc(c)
}

// post перезагрузки пк
func PcReboot(c *gin.Context) {
	spc.RebootPc(c)
}

// post блокировки пк
func PcBlock(c *gin.Context) {
	spc.PcBlock(c)
}

// post разблокировки пк
func PcUnBlock(c *gin.Context) {
	spc.PcUnBlock(c)
}

// get запуска пк старт
func StartPc(c *gin.Context) {
	spc.StartPc(c)
}

// get срикна пк
func ScreenPc(c *gin.Context) {
	spc.Screen(c)
}

// get страницы с пк
func PcDataPage(c *gin.Context) {
	spc.PcPage(c)
}

// post отправки ссылки на пк
func LinkPc(c *gin.Context) {
	spc.PcLink(c)
}

// post хранилища пк
func LSPc(c *gin.Context) {
	spc.PcLS(c)
}

// post редактирования пк
func EditPc(c *gin.Context) {
	spc.Edit(c)
}

// post удаления пк
func DelPC(c *gin.Context) {
	spc.Del(c)
}

func VncPC(c *gin.Context) {
	spc.Vnc(c)
}
