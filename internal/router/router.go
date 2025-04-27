package router

import (
	"database/sql"

	"github.com/alfaa19/kredi-plus-test/internal/handler"
	"github.com/alfaa19/kredi-plus-test/internal/repository"
	"github.com/alfaa19/kredi-plus-test/internal/service"
	locker "github.com/alfaa19/kredi-plus-test/pkg"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, db *sql.DB) {
	// Init Repositories
	userRepo := repository.NewUserRepository(db)
	limitRepo := repository.NewLimitRepository(db)
	trxRepo := repository.NewTransactionRepository(db)

	// Init Services
	locker := locker.NewLocker()
	userService := service.NewUserService(userRepo)
	limitService := service.NewLimitService(limitRepo)
	trxService := service.NewTransactionService(db, limitRepo, trxRepo, locker)

	// Init Handlers
	userHandler := handler.NewUserHandler(userService)
	limitHandler := handler.NewLimitHandler(limitService)
	trxHandler := handler.NewTransactionHandler(trxService)

	// ROUTES

	// User Endpoints
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users/:id", userHandler.GetUserByID)

	// Limit Endpoints
	e.POST("/limits", limitHandler.CreateLimit)
	e.GET("/limits/:user_id", limitHandler.GetLimitsByUserID)

	// Transaction Endpoints
	e.POST("/transactions", trxHandler.CreateTransaction)
	e.GET("/transactions/:id", trxHandler.GetTransactionByID)
	e.GET("/transactions/user/:user_id", trxHandler.GetAllTransactionsByUserID)

}
