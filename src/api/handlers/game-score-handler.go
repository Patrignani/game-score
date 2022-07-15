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

type GameScoreHandlerImp struct {
	insertHandler services.GameScoreInsertHandler
	getAllHandler services.GameScoreGetAllHandler
}

func NewGameScoreHandler(insertHandler services.GameScoreInsertHandler, getHandler services.GameScoreGetAllHandler) ApiFactory {
	return &GameScoreHandlerImp{
		insertHandler,
		getHandler,
	}
}

func (g *GameScoreHandlerImp) insert(echoContext echo.Context) error {
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

func (g *GameScoreHandlerImp) get(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusOK, g.getAllHandler.Get())
}

func (g *GameScoreHandlerImp) Run(e *echo.Echo) {

	b := e.Group("/game-score")
	b.POST("", g.insert)
	b.GET("", g.get)
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
