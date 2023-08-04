package handler

import (
	"example/clean-arch/model"
	"example/clean-arch/model/apperrors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Me(c *gin.Context) {
	user, _ := c.Get("user")
	uid := user.(*model.User).UID

	u, err := h.UserService.Get(c, uid)
	if err != nil {
		msg := fmt.Sprintf("no user with UID %v", uid)
		e := apperrors.NewNotFound(msg)
		c.JSON(e.Status, gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": u,
	})
}
