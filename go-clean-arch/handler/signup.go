package handler

import (
	"example/clean-arch/model"
	"example/clean-arch/model/apperrors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type input struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,gte=6,lte=128" json:"password"`
}

func (h *Handler) Signup(c *gin.Context) {
	var input input

	if ok := bindData(c, &input); !ok {
		return
	}

	user := &model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	err := h.UserService.Signup(c, user)

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
