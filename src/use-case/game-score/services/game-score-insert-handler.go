package services

import (
	"points-game/entities"
	"points-game/helpers"
	"points-game/infra/repository"
	useCase "points-game/use-case/game-score"
	validator "points-game/use-case/game-score/validators"
)

type GameScoreInsertHandler interface {
	Insert(useCase useCase.GameScoreIntserUseCase) []string
}

type GameScoreInsertHandlerImp struct {
	repository repository.GameScoreRepository
}

func NewGameScoreInsertHandler() GameScoreInsertHandler {
	return &GameScoreInsertHandlerImp{repository.NewGameScoreRepository()}
}

func (g GameScoreInsertHandlerImp) Insert(useCase useCase.GameScoreIntserUseCase) []string {
	er := validator.ValidateInsert(useCase)

	if er == nil {
		existRecord := g.repository.ExistRecord(useCase.Point)

		entity, err := entities.NewGameScore(useCase.Point, existRecord == nil)
		er = err

		if er == nil {
			if existRecord == nil {
				g.repository.UpdateRecorrentRecord()
			}

			g.repository.Insert(entity)
		}
	}

	return helpers.GetAllErrors(er)
}
