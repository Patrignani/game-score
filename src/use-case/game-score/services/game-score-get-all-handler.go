package services

import (
	"points-game/boundaries"
	"points-game/infra/repository"
)

type GameScoreGetAllHandler interface {
	Get() boundaries.ResultGameScore
}

type GameScoreGetAllHandlerImp struct {
	repository repository.GameScoreRepository
}

func NewGameScoreGetAllHandler() GameScoreGetAllHandler {
	return &GameScoreGetAllHandlerImp{repository.NewGameScoreRepository()}
}

func (g GameScoreGetAllHandlerImp) Get() boundaries.ResultGameScore {
	entities := g.repository.GetAll()
	var result boundaries.ResultGameScore
	var totalPontos int = 0

	if len(entities) > 0 {

		result = boundaries.ResultGameScore{
			JogosDisputados: len(entities),
			MenorPontuacao:  entities[0].Points,
		}

		for _, value := range entities {
			if value.Record {
				result.QuantidadePontuacaoRecord++
			}

			if value.Points > result.MaiorPontuacao {
				result.MaiorPontuacao = value.Points
			}

			if value.Points < result.MenorPontuacao {
				result.MenorPontuacao = value.Points
			}

			totalPontos += value.Points
		}

		result.MediaPontos = (float32(totalPontos) / float32(result.JogosDisputados))
		result.TotalPontosTemporada = totalPontos
	}

	return result
}
