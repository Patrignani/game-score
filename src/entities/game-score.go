package entities

import (
	"errors"
	msgError "points-game/constants"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameScore struct {
	ID            UniqueEntityID `bson:"_id,json:"id"`
	Points        int            `bson:"points,json:"points"`
	GameDate      time.Time      `bson:"game-date,json:"game-date"`
	Record        bool           `bson:"record,json:"record"`
	CurrentRecord bool           `bson:"current-record,json:"current-record"`
}

func NewGameScore(points int, currentRecord bool) (*GameScore, []error) {
	game := &GameScore{
		ID:            primitive.NewObjectID().String(),
		Points:        points,
		Record:        currentRecord,
		CurrentRecord: currentRecord,
		GameDate:      time.Now(),
	}

	return game, game.Validate()
}

func (g *GameScore) Validate() []error {
	var err []error
	if g.Points < 0 {
		err = append(err, errors.New(msgError.INVALID_POINTS))
	}

	baseValidator := Validate(g.ID)

	if baseValidator != nil {
		err = append(err, baseValidator)
	}

	return err
}
