package server

import (
	"github.com/dbashirov/link-shrinker/internal/shorten"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
	shortener *shorten.Service
}