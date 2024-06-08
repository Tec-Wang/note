package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"wangzheng/brain/internal/domain"
	"wangzheng/brain/internal/entity"
)

type DiaryService struct {
	DiaryDomain *domain.Diary
}

func NewDiaryService(mongodb *mongo.Database, sqlDB *gorm.DB) *DiaryService {
	return &DiaryService{
		DiaryDomain: domain.NewDairy(mongodb, sqlDB),
	}
}

func (s *DiaryService) GetDiary(ctx context.Context, id int64) (*entity.Diary, error) {
	diary, err := s.DiaryDomain.GetDiary(ctx, id)
	if err != nil {
		return &entity.Diary{}, err
	}
	return diary, err
}
