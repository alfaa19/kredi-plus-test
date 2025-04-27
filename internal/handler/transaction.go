package handler

import (
	"net/http"
	"strconv"

	"github.com/alfaa19/kredi-plus-test/internal/dto"
	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/alfaa19/kredi-plus-test/internal/service"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	var req dto.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	trx := model.Transaction{
		UserID:         req.UserID,
		ContractNumber: req.ContractNumber,
		OTRAmount:      req.OTRAmount,
		AdminFee:       req.AdminFee,
		Installment:    req.Installment,
		Interest:       req.Interest,
		AssetName:      req.AssetName,
	}

	if err := h.transactionService.CreateTransaction(&trx, req.TenorMonths); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Transaction created successfully"})
}

func (h *TransactionHandler) GetTransactionByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid transaction id"})
	}

	trx, err := h.transactionService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, trx)
}

func (h *TransactionHandler) GetAllTransactionsByUserID(c echo.Context) error {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	transactions, err := h.transactionService.GetAllByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transactions)
}
