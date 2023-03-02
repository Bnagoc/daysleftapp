package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	service Service
}

func New(service Service) *Endpoint {
	return &Endpoint{
		service: service,
	}
}

func (endpoint *Endpoint) Status(context echo.Context) error {
	days := endpoint.service.DaysLeft()

	daysLeft := fmt.Sprintf("Days left: %d", days)

	if err := context.String(http.StatusOK, daysLeft); err != nil {
		return err
	}

	return nil
}
