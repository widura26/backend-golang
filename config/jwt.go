package config

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Skipper                middleware.Skipper
	BeforeFunc             middleware.BeforeFunc
	SuccessHandler         func(c echo.Context)
	ErrorHandler           func(c echo.Context, err error) error
	ContinueOnIgnoredError bool
	ContextKey             string
	SigningKey             interface{}
	SigningKeys            map[string]interface{}
	SigningMethod          string
	KeyFunc                jwt.Keyfunc
	TokenLookup            string
	TokenLookupFuncs       []middleware.ValuesExtractor
	ParseTokenFunc         func(c echo.Context, auth string) (interface{}, error)
	NewClaimsFunc          func(c echo.Context) jwt.Claims
}
