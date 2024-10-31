package handler

import (
	serrors "github.com/antalkon/zic_server/internal/services/sErrors"
	"github.com/gin-gonic/gin"
)

func NotAuth(c *gin.Context) {
	serrors.Page403(c)
}

func NotFound(c *gin.Context) {
	serrors.Page404(c)
}
