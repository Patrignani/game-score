package entities

import (
	"errors"
	msgError "points-game/constants"
	"strings"
)

type UniqueEntityID = string

func Validate(id UniqueEntityID) error {

	idValidate := strings.Trim(id, " ")
	if idValidate == "" {
		return errors.New(msgError.INVALID_UNIQUE_ENTITY_ID)
	}

	return nil
}
