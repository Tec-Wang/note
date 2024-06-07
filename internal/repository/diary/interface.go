package diary

import "wangzheng/brain/internal/entity"

// Repository interface for diary repository
type Repository interface {
	Save(entity.Diary) error
	GetAll(string) ([]entity.Diary, error)
	Get(string) (*entity.Diary, error)
	Update(entity.Diary) error
	Delete(string) error
}

// NewDiaryRepository creates a new diary repository
func NewDiaryRepository(dbType string) Repository {
	switch dbType {
	case "mongodb":
		return newMongoDBDiaryRepository(GetMongoDB())
	case "mysql":
		return newMySQLDiaryRepository(GetMySQL())
	default:
		return newMySQLDiaryRepository(GetMySQL())
	}
}
