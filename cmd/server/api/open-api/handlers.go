package openapi

import (
	"github.com/gin-gonic/gin"
)

func InitApis(r *gin.Engine) {
	openApi := r.Group("open-api")
	openApi.GET("test", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("test"))
	})
}
