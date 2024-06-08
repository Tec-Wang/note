package diary

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"wangzheng/brain/internal/entity"
)

func GetMongoURI() string {
	return "mongodb://localhost:27017"
}

// MongoAndMysqlRepository implements MongoAndMysqlRepository using MongoDB
type MongoAndMysqlRepository struct {
	mongoDB *mongo.Database
	sqlDB   *gorm.DB
}

func newMongoAndMysqlRepository(mongoDB *mongo.Database, sqlDB *gorm.DB) Repository {
	return &MongoAndMysqlRepository{
		mongoDB: mongoDB,
		sqlDB:   sqlDB,
	}
}

func (r *MongoAndMysqlRepository) Get(ctx context.Context, id int64) (*entity.Diary, error) {
	var entry entity.Diary
	if err := r.sqlDB.First(&entry, id).Error; err != nil {
		return nil, err
	}

	if err := r.mongoDB.Collection("diaries").FindOne(ctx, bson.M{"_id": id}).Decode(&entry); err != nil {
		return nil, err
	}

	return &entry, nil
}
