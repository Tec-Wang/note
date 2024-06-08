package entity

import "gorm.io/gorm"

// Diary represents a single diary entry
type Diary struct {
	gorm.Model
	ID      int64        ``
	UserID  uint         `json:"userID" gorm:"user_id"`
	MongoID string       `json:"mongoID"  gorm:"mongo_id"`
	Content DiaryContent `json:"content" gorm:"-"`
}

func (d *Diary) TableName() string {
	return "diaries"
}

type DiaryContent struct {
	MongoID string `json:"_id" bson:"_id,omitempty"`
	Content string `json:"content" bson:"content,omitempty"`
}
