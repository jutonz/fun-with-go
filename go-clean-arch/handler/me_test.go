package handler

import (
	"encoding/json"
	"example/clean-arch/model"
	"example/clean-arch/model/apperrors"
	"example/clean-arch/model/mocks"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUser := &model.User{
			UID:   uid,
			Email: "me@t.co",
			Name:  "Testing Person",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUser, nil)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{UID: uid})
		})

		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/api/account/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"data": mockUser,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		// assert.Equal(t, respBody, rr.Body.Bytes())
		require.JSONEq(t, string(respBody), rr.Body.String())
		mockUserService.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		uid, _ := uuid.NewRandom()
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, uid).Return(nil, fmt.Errorf("a problem"))

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{UID: uid})
		})

		NewHandler(&Config{
			Router:      router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/api/account/me", nil)
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)

		errorMessage := fmt.Sprintf("no user with UID %v", uid)
		error := apperrors.NewNotFound(errorMessage)
		expectedRespBody, err := json.Marshal(gin.H{"error": error})
		assert.NoError(t, err)

		assert.Equal(t, error.Status, rr.Code)
		// assert.Equal(t, expectedRespBody, rr.Body.Bytes())
		require.JSONEq(t, string(expectedRespBody), rr.Body.String())
	})
}
