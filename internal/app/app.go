package app

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	customMiddleware "github.com/alfaa19/kredi-plus-test/internal/middleware"
	"github.com/alfaa19/kredi-plus-test/internal/router"
)

type App struct {
	Echo *echo.Echo
	DB   *sql.DB
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewApp(db *sql.DB) *App {
	e := echo.New()

	// Setup Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Setup Custom Error Handler
	e.HTTPErrorHandler = customMiddleware.CustomHTTPErrorHandler

	// Middleware Basic
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS()) // kalau mau open cors

	// Setup Router
	router.NewRouter(e, db)

	return &App{
		Echo: e,
		DB:   db,
	}
}

func (a *App) Run(addr string) error {
	return a.Echo.Start(addr)
}
