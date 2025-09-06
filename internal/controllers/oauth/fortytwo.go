package oauth

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"server/internal/controllers/auth"
	"server/internal/services"
	oauthService "server/internal/services/oauth2"
	"server/internal/services/users"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func FortyTwo(c echo.Context) error {
	config := oauthService.Providers()["42"]
	if config == nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "42 API method not implemented",
		})
	}
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusFound, url)
}

type OAuth2FortyTwoRedirect struct {
	Code string `validate:"required" query:"code"`
}

func FortyTwoCallback(c echo.Context) error {
	config := oauthService.Providers()["42"]
	if config == nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "42 API method not implemented",
		})
	}

	var body OAuth2FortyTwoRedirect

	err := c.Bind(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	token, err := config.Exchange(context.Background(), body.Code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "Unable to retrive token: " + err.Error(),
		})
	}

	err = services.ValidateStruct(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	req, err := http.NewRequest("GET", "https://api.intra.42.fr/v2/me", nil)
	if err != nil {
		services.Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	token.SetAuthHeader(req)
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		services.Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if res.StatusCode != http.StatusOK {
		services.Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var resBody bytes.Buffer

	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return err
	}

	var APIUserRes map[string]interface{}
	if err := json.Unmarshal(resBody.Bytes(), &APIUserRes); err != nil {
		return err
	}

	provideIdRaw, ok := APIUserRes["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Could not get user_id from method",
		})
	}
	provideId := strconv.Itoa(int(provideIdRaw))
	user, err := users.GetUserByProviderId("42", provideId)
	if err == gorm.ErrRecordNotFound {
		newUser := users.CreateUserType{
			Provider:   "42",
			ProviderId: provideId,
			Email:      APIUserRes["email"].(string),
			Name:       APIUserRes["displayname"].(string),
		}
		createdUser, err := users.CreateUser(newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": err.Error(),
			})
		}
		user = createdUser
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	response, err := auth.RevokeToken(user, "")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
