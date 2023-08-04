package handler

import (
	"example/clean-arch/model/apperrors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func bindData(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %v\n", err)

		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			error := apperrors.NewBadRequest()

			c.JSON(error.Status, gin.H{
				"error":       error,
				"invalidArgs": invalidArgs,
			})
		} else {
			// err not an instance of validator.ValidationErrors
			error := apperrors.NewInternalError()
			c.JSON(error.Status, gin.H{"error": error})
		}

		return false
	}

	return true
}
