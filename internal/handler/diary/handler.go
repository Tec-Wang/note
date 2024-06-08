package diary

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"wangzheng/brain/internal/service"
)

type Handler struct {
	r            *gin.RouterGroup
	diaryService *service.DiaryService
}

func NewDiaryHandler(r *gin.RouterGroup, sqlDB *gorm.DB, mongoDB *mongo.Database) *Handler {
	return &Handler{
		r:            r,
		diaryService: service.NewDiaryService(mongoDB, sqlDB),
	}
}

func (h *Handler) Handle() {
	h.r.GET("/:id", h.HandleGetDiary)
}

func (h *Handler) HandleGetDiary(ctx *gin.Context) {
	var id idParam
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	byID, err := h.diaryService.GetDiary(ctx, id.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
		"data":    byID,
	})
}
