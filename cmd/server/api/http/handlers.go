package http

import (
	"wangzheng/brain/cmd/server/api/http/brain"

	"github.com/gin-gonic/gin"
)

func InitApis(r *gin.Engine) {
	api := r.Group("api")
	brain.InitApis(api)
}
