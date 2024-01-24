package usecase

import (
	"errors"
	"fmt"

	"github.com/Adilfarooque/Footsgo_Ecommerce/helper"
	"github.com/Adilfarooque/Footsgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footsgo_Ecommerce/utils/models"
)

/*
The function to handle user registration by checking for
existing users with the same email or phone,
hashing the password, creating a referral code,
and generating access and refresh tokens upon successful registration
*/
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
	if phone != nil {
		return &models.TokenUser{}, errors.New("user with this phone is already exists")
	}

	hashPassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return &models.TokenUser{}, errors.New("error in hashing password")
	}
	user.Password = hashPassword

	return
}
