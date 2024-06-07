package api

import (
	"github.com/gin-gonic/gin"
	"wangzheng/brain/internal/handler"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	handler.NewDiaryHandler(v1.Group("/diary")).Handle()
}
