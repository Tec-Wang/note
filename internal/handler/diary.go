package handler

import (
	"github.com/gin-gonic/gin"
	"wangzheng/brain/internal/service"
)

type DiaryHandler struct {
	r            *gin.RouterGroup
	diaryService *DiaryService
}

func NewDiaryHandler(r *gin.RouterGroup) *DiaryHandler {
	return &DiaryHandler{
		r:            r,
		diaryService: service.NewDiaryService(),
	}
}

func (h *DiaryHandler) Handle() {
	h.r.GET("", HandleGetDiary)
}

func HandleGetDiary(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
