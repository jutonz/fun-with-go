package handler

import (
	"bytes"
	"encoding/json"
	"example/clean-arch/model"
	"example/clean-arch/model/apperrors"
	"example/clean-arch/model/mocks"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		user := &model.User{
			Email:    "me@t.co",
			Password: "password123",
		}
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    user.Email,
			"password": user.Password,
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("content-type", "application/json")

		router.ServeHTTP(rr, request)

		expectedResp, err := json.Marshal(gin.H{"data": user})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rr.Code)
		require.JSONEq(t, string(expectedResp), rr.Body.String())
		mockUserService.AssertExpectations(t)
	})

	t.Run("UserService returns generic error", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(fmt.Errorf("an error"))

		rr := httptest.NewRecorder()
		router := gin.Default()
		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "email@t.co",
			"password": "password123",
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("content-type", "application/json")

		router.ServeHTTP(rr, request)

		expectedError := apperrors.NewInternalError()
		expectedResp, err := json.Marshal(gin.H{
			"error": expectedError,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		require.JSONEq(t, string(expectedResp), rr.Body.String())
		mockUserService.AssertExpectations(t)
	})

	t.Run("Email and password required", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("content-type", "application/json")

		router.ServeHTTP(rr, request)

		expectedError := apperrors.NewBadRequest()
		expectedResp, err := json.Marshal(gin.H{
			"error": expectedError,
			"invalidArgs": []gin.H{
				{
					"field": "Email",
					"value": "",
					"tag":   "required",
					"param": "",
				},
				{
					"field": "Password",
					"value": "",
					"tag":   "required",
					"param": "",
				},
			},
		})
		log.Printf("%v", string(expectedResp))
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		require.JSONEq(t, string(expectedResp), rr.Body.String())
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Password is too short", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "valid@email.com",
			"password": "12345",
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("content-type", "application/json")

		router.ServeHTTP(rr, request)

		expectedError := apperrors.NewBadRequest()
		expectedResp, err := json.Marshal(gin.H{
			"error": expectedError,
			"invalidArgs": []gin.H{
				{
					"field": "Password",
					"value": "12345",
					"tag":   "gte",
					"param": "6",
				},
			},
		})
		log.Printf("%v", string(expectedResp))
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		require.JSONEq(t, string(expectedResp), rr.Body.String())
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Email is invalid format", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "bad email",
			"password": "avalidpassword",
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("content-type", "application/json")

		router.ServeHTTP(rr, request)

		expectedError := apperrors.NewBadRequest()
		expectedResp, err := json.Marshal(gin.H{
			"error": expectedError,
			"invalidArgs": []gin.H{
				{
					"field": "Email",
					"value": "bad email",
					"tag":   "email",
					"param": "",
				},
			},
		})
		log.Printf("%v", string(expectedResp))
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		require.JSONEq(t, string(expectedResp), rr.Body.String())
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Wrong content type", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "bad email",
			"password": "avalidpassword",
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("content-type", "application/xml")

		router.ServeHTTP(rr, request)

		expectedError := apperrors.NewInternalError()
		expectedResp, err := json.Marshal(gin.H{
			"error": expectedError,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		require.JSONEq(t, string(expectedResp), rr.Body.String())
		mockUserService.AssertNotCalled(t, "Signup")
	})
}
