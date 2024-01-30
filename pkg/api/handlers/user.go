package handlers

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/utils/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func (ur *UserHandler)UserSignup(c *gin.Context){
	var SignupDetails models.UserSignUp
	if err := c.ShouldBindJSON(&SignupDetails);err != nil{
		errs := response.
	}
}