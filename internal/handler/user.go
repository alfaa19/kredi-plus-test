package handler

import (
	"net/http"
	"strconv"

	"github.com/alfaa19/kredi-plus-test/internal/dto"
	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/alfaa19/kredi-plus-test/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req dto.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	user := model.User{
		NIK:            req.NIK,
		FullName:       req.FullName,
		LegalName:      req.LegalName,
		PlaceOfBirth:   req.PlaceOfBirth,
		DateOfBirth:    req.DateOfBirth,
		Salary:         req.Salary,
		KTPPhotoURL:    req.KTPPhotoURL,
		SelfiePhotoURL: req.SelfiePhotoURL,
	}

	if err := h.userService.Create(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User created successfully"})
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
