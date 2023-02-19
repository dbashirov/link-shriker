package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dbashirov/link-shrinker/internal/model"
	"github.com/dbashirov/link-shrinker/internal/shorten"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type shortener interface {
	Shorten(context.Context, model.ShortenInput) (*model.Shortening, error)
}

type shortenRequest struct {
	URL        string `json:"url" validate:"required,url"`
	Identifier string `json:"identifier,omitempty" validate:"omitempty,alphanum"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url,omitempty"`
	Message  string `json:"message,omitempty"`
}

func HandleShorten(shortener shortener) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req shortenRequest
		if err := c.Bind(&req); err != nil {
			return err
		}

		if err := c.Validate(req); err != nil {
			return err
		}

		userToken, ok := c.Get("user").(*jwt.Token)
		if !ok {
			log.Println("error: user is not presented in context")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		userClaims, ok := userToken.Claims.(*model.UserClaims)
		if !ok {
			log.Println("error: failed to get user claims from token")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		var identifier string
		if strings.TrimSpace(req.Identifier) != "" {
			identifier = req.Identifier
		}

		input := model.ShortenInput{
			RawURL:     req.URL,
			Identifier: identifier,
			CreatedBy:  userClaims.User.GithubLogin,
		}

		shortening, err := shortener.Shorten(c.Request().Context(), input)
		if err != nil {
			if errors.Is(err, model.ErrIdentifierExists) {
				return echo.NewHTTPError(http.StatusConflict, model.ErrIdentifierExists.Error())
			}
			log.Printf("error shortening url %q: %v", req.URL, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		shortURL, err := shorten.PrependBaseURL(config.Get().BaseURL)
	}
}
