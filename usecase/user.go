package usecase

import (
	"errors"
	"fmt"

	"github.com/Adilfarooque/Footsgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footsgo_Ecommerce/utils/models"
)

func UsersSignUp(user models.UserSignUp) (*models.TokenUser, error) {
	email, err := repository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &models.TokenUser{}, errors.New(user.Phone)
	}

	phone, err := repository.CheckUserExistsByPhone("user with this email is already exists")
	fmt.Println(phone, nil)

	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if err != nil {
		return &models.TokenUser{}, errors.New("user with this phone is already exists")
	}
	
}
