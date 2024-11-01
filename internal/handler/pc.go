package handler

import (
	spc "github.com/antalkon/zic_server/internal/services/sPc"
	"github.com/gin-gonic/gin"
)

func AddNewPc(c *gin.Context) {
	spc.AddPc(c)
}

func PcCount(c *gin.Context) {
	spc.PcCount(c)
}

func ServerPing(c *gin.Context) {
	spc.PcPing(c)
}

func PcOff(c *gin.Context) {
	spc.OffPc(c)
}

func PcReboot(c *gin.Context) {
	spc.RebootPc(c)
}

func PcBlock(c *gin.Context) {
	spc.PcBlock(c)
}

func PcUnBlock(c *gin.Context) {
	spc.PcUnBlock(c)
}

func StartPc(c *gin.Context) {
	spc.StartPc(c)
}

func ScreenPc(c *gin.Context) {
	spc.Screen(c)
}

func PcDataPage(c *gin.Context) {
	spc.PcPage(c)
}

func LinkPc(c *gin.Context) {
	spc.PcLink(c)
}

func LSPc(c *gin.Context) {
	spc.PcLS(c)
}

func EditPc(c *gin.Context) {
	spc.Edit(c)
}

func DelPC(c *gin.Context) {
	spc.Del(c)
}
