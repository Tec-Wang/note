package entity

import "time"

// DiaryEntry represents a single diary entry
type DiaryEntry struct {
	ID        string    `bson:"_id"`
	UserID    string    `bson:"user_id"`
	Content   string    `bson:"content"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
