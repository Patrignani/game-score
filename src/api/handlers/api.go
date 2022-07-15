package handlers

import "github.com/labstack/echo"

type ApiFactory interface {
	Run(e *echo.Echo)
}
