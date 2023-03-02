package app

import (
	"fmt"
	"log"

	"github.com/Bnagoc/daysapp/internal/app/endpoint"
	"github.com/Bnagoc/daysapp/internal/app/service"
	"github.com/labstack/echo/v4"

	"github.com/Bnagoc/daysapp/internal/app/mw"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	app.service = service.New()
	app.endpoint = endpoint.New(app.service)

	app.echo = echo.New()
	app.echo.Use(mw.CheckRole)

	app.echo.GET("/status", app.endpoint.Status)

	return app, nil

}

func (app *App) Run() error {
	fmt.Println("server running...")

	if err := app.echo.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	return nil
}
