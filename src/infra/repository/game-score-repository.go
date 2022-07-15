package repository

import (
	"context"
	"points-game/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameScoreRepository interface {
	UpdateRecorrentRecord()
	ExistRecord(point int) *entities.GameScore
	Insert(gameScore *entities.GameScore)
	GetAll() []entities.GameScore
}

type GameScoreRepositoryImp struct {
	collection *mongo.Collection
}

func NewGameScoreRepository(collection *mongo.Collection) GameScoreRepository {
	return &GameScoreRepositoryImp{collection: collection}
}

func (g GameScoreRepositoryImp) Insert(gameScore *entities.GameScore) {
	g.collection.InsertOne(context.TODO(), gameScore)
}

func (g GameScoreRepositoryImp) UpdateRecorrentRecord() {
	filter := bson.D{{"current-record", true}}
	update := bson.D{{"$set", bson.D{{"current-record", false}}}}

	g.collection.UpdateMany(context.TODO(), filter, update)
}

func (g GameScoreRepositoryImp) ExistRecord(point int) *entities.GameScore {

	var result *entities.GameScore

	filter := bson.M{"points": bson.M{"$gte": point}}

	g.collection.FindOne(context.TODO(), filter).Decode(&result)

	return result
}

func (g GameScoreRepositoryImp) GetAll() []entities.GameScore {
	var result []entities.GameScore
	cursor, _ := g.collection.Find(context.TODO(), bson.D{})
	cursor.All(context.TODO(), &result)
	return result
}
