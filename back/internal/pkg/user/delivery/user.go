package delivery

import (
	"github.com/labstack/echo/v4"
)

type UserDeliveryModel struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	BIO      string `json:"bio,omitempty"`
}

type UserDelivery interface {
	Authenticate(c echo.Context) error
	Register(c echo.Context) error
}
