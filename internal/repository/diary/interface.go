package diary

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"wangzheng/brain/internal/entity"
)

// Repository interface for diary repository
type Repository interface {
	Get(ctx context.Context, int64 int64) (*entity.Diary, error)
}

// NewDiaryRepository creates a new diary repository
func NewDiaryRepository(mongoDB *mongo.Database, mysqlDB *gorm.DB) Repository {
	return newMongoAndMysqlRepository(mongoDB, mysqlDB)
}
