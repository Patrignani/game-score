package ioc

import (
	"points-game/api/handlers"
	"points-game/infra"
	"points-game/infra/repository"
	"points-game/use-case/game-score/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type Ioc struct{}

func NewIoc() Ioc {
	return Ioc{}
}

func (i Ioc) GetGameScore() handlers.ApiFactory {
	var collection *mongo.Collection = infra.NewMongoContext("mongodb://localhost:27017").GetCollection("game", "game-score")
	repository := repository.NewGameScoreRepository(collection)
	getAllHandler := services.NewGameScoreGetAllHandler(repository)
	insertHandler := services.NewGameScoreInsertHandler(repository)
	api := handlers.NewGameScoreHandler(insertHandler, getAllHandler)
	return api
}
