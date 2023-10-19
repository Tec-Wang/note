package api

import (
	"wangzheng/brain/cmd/server/api/http"
	openapi "wangzheng/brain/cmd/server/api/open-api"

	"github.com/gin-gonic/gin"
)

func InitApis(r *gin.Engine) {
	openapi.InitApis(r)
	http.InitApis(r)
}
