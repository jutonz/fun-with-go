package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PingGet() gin.HandlerFunc {
  return func(c *gin.Context) {
    msg := fmt.Sprintf("value from env is %v", os.Getenv("VALUE"))

    c.IndentedJSON(http.StatusOK, map[string]string{
      "message": msg,
    })
  }
}
