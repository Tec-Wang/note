package service

import "wangzheng/brain/internal/domain"

type DiaryService struct {
	DiaryDomain *domain.Diary
}

func NewDiaryService() *DiaryService {
	return &DiaryService{
		DiaryDomain: domain.NewDairy(),
	}
}

func (s *DiaryService) GetDiaryByID(id int) (domain.Diary, error) {

}
