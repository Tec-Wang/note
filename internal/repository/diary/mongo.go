package diary

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"wangzheng/brain/internal/entity"
)

func GetMongoURI() string {
	return "mongodb://localhost:27017"
}

// MongoDBDiaryRepository implements Repository using MongoDB
type MongoDBDiaryRepository struct {
	db *mongo.Database
}

func newMongoDBDiaryRepository(db *mongo.Database) *MongoDBDiaryRepository {
	return &MongoDBDiaryRepository{db: db}
}

func (r *MongoDBDiaryRepository) Save(entry entity.Diary) error {
	_, err := r.db.Collection("diaries").InsertOne(context.TODO(), entry)
	return err
}

func (r *MongoDBDiaryRepository) GetAll(userID string) ([]entity.Diary, error) {
	var entries []entity.Diary
	cursor, err := r.db.Collection("diaries").Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &entries)
	return entries, err
}

func (r *MongoDBDiaryRepository) Get(id string) (*entity.Diary, error) {
	var entry entity.Diary
	err := r.db.Collection("diaries").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&entry)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *MongoDBDiaryRepository) Update(entry entity.Diary) error {
	_, err := r.db.Collection("diaries").UpdateOne(context.TODO(), bson.M{"_id": entry.ID}, bson.M{"$set": entry})
	return err
}

func (r *MongoDBDiaryRepository) Delete(id string) error {
	_, err := r.db.Collection("diaries").DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func GetMongoDB() *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(GetMongoURI()))
	if err != nil {
		panic(err)
	}
	return client.Database("diary")
}
