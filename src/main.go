package main

import (
	"context"
	"points-game/api/handlers"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	gamescoreHandler := handlers.NewGameScoreHandler()

	b := e.Group("/game-score")
	b.POST("", gamescoreHandler.Insert)
	b.GET("", gamescoreHandler.Get)
	e.Logger.Fatal(e.Start("localhost:9098"))
}
