package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	useCase "points-game/use-case/game-score"
	"points-game/use-case/game-score/services"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type GameScoreHandler interface {
	Insert(echoContext echo.Context) error
	Get(echoContext echo.Context) error
}

type GameScoreHandlerImp struct {
	insertHandler services.GameScoreInsertHandler
	getAllHandler services.GameScoreGetAllHandler
}

func NewGameScoreHandler() GameScoreHandler {
	return &GameScoreHandlerImp{
		services.NewGameScoreInsertHandler(),
		services.NewGameScoreGetAllHandler(),
	}
}

func (g *GameScoreHandlerImp) Insert(echoContext echo.Context) error {
	insertUseCase := useCase.GameScoreIntserUseCase{}

	er := getBody(echoContext, &insertUseCase)

	if er != nil {
		return er
	}

	errs := g.insertHandler.Insert(insertUseCase)

	if errs != nil {
		return echoContext.JSON(http.StatusBadRequest, errs)
	}

	return echoContext.JSON(http.StatusOK, "")
}

func (g *GameScoreHandlerImp) Get(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusOK, g.getAllHandler.Get())
}

func getBody[T any](echoContext echo.Context, object *T) error {
	body, err := ioutil.ReadAll(echoContext.Request().Body)

	if err != nil {
		log.Error("[BackstageHandler UserUpdate] Error reading body", err)
		return echoContext.JSON(http.StatusBadRequest, nil)
	}

	err = json.Unmarshal(body, object)

	if err != nil {
		log.Error("[BackstageHandler UserUpdate] Error reading body", err)
		return echoContext.JSON(http.StatusBadRequest, nil)
	}

	return nil
}
