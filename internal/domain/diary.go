package domain

import (
	"wangzheng/brain/internal/entity"
	model "wangzheng/brain/internal/model/diary"
)

// DiaryRepository interface for diary repository
type DiaryRepository interface {
	Save(entity.DiaryEntry) error
	GetAll(string) ([]entity.DiaryEntry, error)
	Get(string) (*entity.DiaryEntry, error)
	Update(entity.DiaryEntry) error
	Delete(string) error
}

// NewDiaryRepository creates a new diary repository
func NewDiaryRepository(dbType string) DiaryRepository {
	switch dbType {
	case "mongodb":
		return model.NewMongoDBDiaryRepository(model.GetMongoDB())
	case "mysql":
		return model.NewMySQLDiaryRepository(model.GetMySQL())
	default:
		return nil
	}
}
