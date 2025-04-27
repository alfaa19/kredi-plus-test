package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockTransactionService struct{}

func (m *MockTransactionService) CreateTransaction(trx *model.Transaction, tenorMonths int) error {
	if trx.UserID == 99 {
		return errors.New("failed create transaction")
	}
	return nil
}
func (m *MockTransactionService) GetByID(id int64) (*model.Transaction, error) {
	return &model.Transaction{}, nil
}
func (m *MockTransactionService) GetAllByUserID(userID int64) ([]model.Transaction, error) {
	return []model.Transaction{}, nil
}

func TestCreateTransactionHandler_Success(t *testing.T) {
	e := echo.New()
	handler := NewTransactionHandler(&MockTransactionService{})

	reqBody := map[string]interface{}{
		"user_id":         1,
		"tenor_months":    6,
		"contract_number": "CTR001",
		"otr_amount":      150000,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CreateTransaction(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateTransactionHandler_Fail(t *testing.T) {
	e := echo.New()
	handler := NewTransactionHandler(&MockTransactionService{})

	reqBody := map[string]interface{}{
		"user_id":         99,
		"tenor_months":    6,
		"contract_number": "CTR001",
		"otr_amount":      150000,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CreateTransaction(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
