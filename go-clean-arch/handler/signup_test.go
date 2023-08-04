package handler

import (
	"bytes"
	"encoding/json"
	"example/clean-arch/model/apperrors"
	"example/clean-arch/model/mocks"
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
		
	})

	t.Run("UserService returns generic error", func(t *testing.T) {
		
	})

	t.Run("Email and password required", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		NewHandler(&Config{
			Router: router,
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
					"tag": "required",
					"param": "",
				},
				{
					"field": "Password",
					"value": "",
					"tag": "required",
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
		
	})

	t.Run("Email is invalid format", func(t *testing.T) {
		
	})
}
