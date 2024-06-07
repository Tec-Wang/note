package domain

import (
	"wangzheng/brain/internal/entity"
	"wangzheng/brain/internal/repository/diary"
)

type Diary struct {
	diaryRepository diary.Repository
}

func NewDairy() *Diary {
	return &Diary{
		diaryRepository: diary.NewDiaryRepository("mongo"),
	}
}

func (d *Diary) GetDiary(id string) (*entity.Diary, error) {
	return d.diaryRepository.Get(id)
}
