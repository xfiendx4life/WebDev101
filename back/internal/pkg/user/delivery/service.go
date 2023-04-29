package delivery

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/xfiendx4life/web101/meet/internal/pkg/user/usecase"

	// echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/golang-jwt/jwt/v4"
)

var secret = "newsecret"

type usrDlvr struct {
	usc usecase.UserUsecase
}

type jwtCustomClaims struct {
	ID string `json:"name"`
	jwt.RegisteredClaims
}

func createJWT(id string) (string, error) {
	claims := &jwtCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func New(us usecase.UserUsecase) UserDelivery {
	return &usrDlvr{
		usc: us,
	}
}

func (d *usrDlvr) Authenticate(c echo.Context) error {
	data := c.Request().Body
	var userData []byte
	_, err := data.Read(userData)
	if err != nil {
		c.Logger().Errorf("can't read request body: %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "{error: Bad Request}")
	}
	delModel := UserDeliveryModel{}
	err = json.Unmarshal(userData, &delModel)
	if err != nil {
		c.Logger().Errorf("can't parse from json to model: %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "{error: Bad Request}")
	}
	id, err := d.usc.Authenticate(c.Request().Context(), delModel.Name, delModel.Password)
	if err != nil {
		if err == usecase.ErrorAuthentication {
			c.Logger().Debug("Authintication failed")
			return c.JSON(http.StatusOK, map[string]string{"status": "not authenticate"})
		}
		c.Logger().Errorf("can't parse from json to model: %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "{error: Bad Request}")
	}
	t, err := createJWT(id.String())
	if err != nil {
		c.Logger().Error("can't create token: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)

	}
	return c.JSON(http.StatusOK, echo.Map{"token": t})

}
func (d *usrDlvr) Register(c echo.Context) error {
	return nil
}
