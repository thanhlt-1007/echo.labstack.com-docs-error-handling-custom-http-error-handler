package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func customHTTPErrorHandler(err error, context echo.Context) {
	if context.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	httpError, ok := err.(*echo.HTTPError)
	if ok {
		code = httpError.Code
	}

	context.Logger().Error(httpError)
	context.JSON(code, err)
}

func getPingHanler(context echo.Context) error {
	status := context.QueryParam("status")
	if status != "OK" {
		return echo.ErrBadRequest
	}

	return context.String(
		http.StatusOK,
		"OK",
	)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/ping", getPingHanler)
	e.Logger.Fatal(e.Start(":1323"))
}
