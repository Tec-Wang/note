package entity

import (
	"gorm.io/gorm"
)

// Diary represents a single diary entry
type Diary struct {
	gorm.Model

	MongoID string `bson:"_id" gorm:"m_id" json:"-"`
	UserID  string `bson:"user_id" gorm:"user_id" json:"userID"`
	Content string `bson:"content" gorm:"content" json:"content"`
}
