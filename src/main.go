package main

import (
	"context"
	"points-game/ioc"
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

	ioc.NewIoc().GetGameScore().Run(e)

	e.Logger.Fatal(e.Start("localhost:9098"))
}
