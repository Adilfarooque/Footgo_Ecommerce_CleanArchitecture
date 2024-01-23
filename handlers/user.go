package handlers

import (
	"net/http"

	"github.com/Adilfarooque/Footsgo_Ecommerce/utils/models"
	"github.com/Adilfarooque/Footsgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//@Summary User Signup
//@Description user ca signup by giving their details
//@Tags   		User
//@Accept 		json
//Produce 		json
//Param   		signup by modles.UserSignup true "signup"
//Success   200	{object} response.Response{}
//Failure   500	{object} response.Response{}
//Routes        /user/signup [POST]
func UserSignup(c *gin.Context){
	//Reading JSON input
	var SignupDetails models.UserSignUp
	if err := c.ShouldBindJSON(&SignupDetails); err != nil{
		errs := response.ClientResponse(http.StatusBadRequest,"Details not in correct format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errs)
		return
	}
	//Data Validation
	
}