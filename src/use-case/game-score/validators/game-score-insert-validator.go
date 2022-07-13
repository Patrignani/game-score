package validators

import (
	"errors"
	msgError "points-game/constants"
	useCase "points-game/use-case/game-score"
)

func ValidateInsert(gameScore useCase.GameScoreIntserUseCase) []error {
	var err []error

	if gameScore.Point < 0 {
		err = append(err, errors.New(msgError.INVALID_POINTS))
	}

	return err
}
