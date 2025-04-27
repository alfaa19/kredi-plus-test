package handler

import (
	"net/http"
	"strconv"

	"github.com/alfaa19/kredi-plus-test/internal/dto"
	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/alfaa19/kredi-plus-test/internal/service"
	"github.com/labstack/echo/v4"
)

type LimitHandler struct {
	limitService service.LimitService
}

func NewLimitHandler(limitService service.LimitService) *LimitHandler {
	return &LimitHandler{limitService: limitService}
}

func (h *LimitHandler) CreateLimit(c echo.Context) error {
	var req dto.LimitRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	limit := model.Limit{
		UserID:      req.UserID,
		TenorMonths: req.TenorMonths,
		LimitAmount: req.LimitAmount,
	}

	if err := h.limitService.Create(&limit); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Limit created successfully"})
}

func (h *LimitHandler) GetLimitsByUserID(c echo.Context) error {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	limits, err := h.limitService.GetAllByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, limits)
}
