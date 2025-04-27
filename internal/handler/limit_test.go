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

type MockLimitService struct{}

func (m *MockLimitService) Create(limit *model.Limit) error {
	if limit.UserID == 99 {
		return errors.New("failed create limit")
	}
	return nil
}
func (m *MockLimitService) GetByUserIDAndTenor(userID int64, tenor int) (*model.Limit, error) {
	return &model.Limit{}, nil
}
func (m *MockLimitService) GetAllByUserID(userID int64) ([]model.Limit, error) {
	return []model.Limit{}, nil
}
func (m *MockLimitService) Update(limit *model.Limit) error { return nil }

func TestCreateLimitHandler_Success(t *testing.T) {
	e := echo.New()
	handler := NewLimitHandler(&MockLimitService{})

	reqBody := map[string]interface{}{
		"user_id":      1,
		"tenor_months": 6,
		"limit_amount": 1000000,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/limits", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CreateLimit(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateLimitHandler_Fail(t *testing.T) {
	e := echo.New()
	handler := NewLimitHandler(&MockLimitService{})

	reqBody := map[string]interface{}{
		"user_id":      99,
		"tenor_months": 6,
		"limit_amount": 1000000,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/limits", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CreateLimit(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetLimitsByUserIDHandler_Success(t *testing.T) {
	e := echo.New()
	handler := NewLimitHandler(&MockLimitService{})

	req := httptest.NewRequest(http.MethodGet, "/limits/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("user_id")
	c.SetParamValues("1")

	err := handler.GetLimitsByUserID(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
