package brain

import (
	"github.com/gin-gonic/gin"
)

func InitApis(r *gin.RouterGroup) {
	brain := r.Group("brain")

	brain.GET("page", Page)
}

func Page(ctx *gin.Context) {
	ctx.Writer.Write([]byte("page"))
}
