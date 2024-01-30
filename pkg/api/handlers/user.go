package handlers

import (
	"net/http"

	services "github.com/Adilfarooque/Footgo_Ecommerce/pkg/usercase/interfaces"
	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/utils/models"
	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	UserUseCase services.UserUseCase
}

func NewUserHandler(useCase services.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: useCase,
	}
}

// Summary User Signup
// Description user can signup by giving their details
// Tags User
// Accept json
// Produce json
// Param  singup body models.UserSignUp true "signup"
// Success 200 {object} response.Response{}
// Faliure 500 {object} response.Response{}
// Router 	/user/signup [POST]
func (ur *UserHandler) UserSignup(c *gin.Context) {
	var SignupDetails models.UserSignUp
	if err := c.ShouldBindJSON(&SignupDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct foramt", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	err := validator.New().Struct(SignupDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := ur.UserUseCase.UsersSignUP(SignupDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	success := response.ClientResponse(http.StatusCreated, "User successful signed up", user, nil)
	c.JSON(http.StatusCreated, success)
}
