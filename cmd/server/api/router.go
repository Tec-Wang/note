package api

import (
	"github.com/gin-gonic/gin"
	"wangzheng/brain/cmd"
	"wangzheng/brain/internal/handler/diary"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	mysqlDB := cmd.Mysql
	mongoClient := cmd.Mongo

	diary.NewDiaryHandler(v1.Group("/diary"), mysqlDB, mongoClient.Database("note")).Handle()
}
