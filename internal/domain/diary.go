package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"wangzheng/brain/internal/entity"
	"wangzheng/brain/internal/repository/diary"
)

type Diary struct {
	diaryRepository diary.Repository
}

func NewDairy(mongoDB *mongo.Database, sqlDB *gorm.DB) *Diary {
	return &Diary{
		diaryRepository: diary.NewDiaryRepository(mongoDB, sqlDB),
	}
}

func (d *Diary) GetDiary(ctx context.Context, id int64) (*entity.Diary, error) {
	return d.diaryRepository.Get(ctx, id)
}
