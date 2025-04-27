package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alfaa19/kredi-plus-test/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct{}

func (m *MockUserService) Create(user *model.User) error {
	if user.FullName == "fail" {
		return errors.New("failed create user")
	}
	return nil
}
func (m *MockUserService) GetByID(id int64) (*model.User, error) {
	if id == 99 {
		return nil, errors.New("user not found")
	}
	return &model.User{
		ID:          id,
		FullName:    "Budi Test",
		NIK:         "1234567890123456",
		DateOfBirth: time.Now(),
	}, nil
}
func (m *MockUserService) Update(user *model.User) error { return nil }
func (m *MockUserService) Delete(id int64) error         { return nil }

func TestCreateUserHandler_Success(t *testing.T) {
	e := echo.New()
	handler := NewUserHandler(&MockUserService{})

	reqBody := map[string]interface{}{
		"nik":            "1234567890123456",
		"full_name":      "Budi Test",
		"legal_name":     "Budi Legal",
		"place_of_birth": "Jakarta",
		"date_of_birth":  "2000-01-01T00:00:00Z",
		"salary":         10000000,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CreateUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateUserHandler_Fail(t *testing.T) {
	e := echo.New()
	handler := NewUserHandler(&MockUserService{})

	reqBody := map[string]interface{}{
		"nik":            "1234567890123456",
		"full_name":      "fail",
		"legal_name":     "fail",
		"place_of_birth": "fail",
		"date_of_birth":  "2000-01-01T00:00:00Z",
		"salary":         10000000,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CreateUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetUserByIDHandler_Success(t *testing.T) {
	e := echo.New()
	handler := NewUserHandler(&MockUserService{})

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := handler.GetUserByID(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUserByIDHandler_Fail(t *testing.T) {
	e := echo.New()
	handler := NewUserHandler(&MockUserService{})

	req := httptest.NewRequest(http.MethodGet, "/users/99", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("99")

	err := handler.GetUserByID(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
