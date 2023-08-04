package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"example/api-testin-v2/main"

	"github.com/gin-gonic/gin"
)

func SetupTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}

func TestGetBooksEmpty(t *testing.T) {
	SetupRouter()
}
