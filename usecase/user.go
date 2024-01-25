package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Adilfarooque/Footsgo_Ecommerce/helper"
	"github.com/Adilfarooque/Footsgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footsgo_Ecommerce/utils/models"
	"github.com/google/uuid"
)

/*
The function to handle user registration by checking for
existing users with the same email or phone,
hashing the password, creating a referral code,
and generating access and refresh tokens upon successful registration
*/
func UsersSignUp(user models.UserSignUp) (*models.TokenUser, error) {
	//Check if the user already exists by email
	email, err := repository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &models.TokenUser{}, errors.New(user.Phone)
	}
	//Check if the user already exists by phone
	phone, err := repository.CheckUserExistsByPhone("user with this email is already exists")
	fmt.Println(phone, nil)

	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if phone != nil {
		return &models.TokenUser{}, errors.New("user with this phone is already exists")
	}
	//Hash the user's password
	hashPassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return &models.TokenUser{}, errors.New("error in hashing password")
	}
	user.Password = hashPassword

	//User signup in the repository
	userData, err := repository.UserSignUp(user)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not add the user")
	}

	//Create a referral code for the user and store it in the database
	//Generate User Refferal
	id := uuid.New().ID()
	//convert UUID into string
	str := strconv.Itoa(int(id))
	//Takes the first 8 character of the string to create refferal code
	userReferral := str[:8]
	/*Create referral entry
	Calls a repository function to create 
	a referral entry for the user in the database, 
	associating the generated referral code with the user
	*/
	err = repository.CreateReferralEntry(userData, userReferral)
	if err != nil {
		return &models.TokenUser{}, err
	}

	if user.ReferralCode != ""{
		//First check whether if a user with that refferalCode exist
		referredUserId, err := repository.GetUserIdFromReferrals(user.ReferralCode)
		if err != nil{
			return &models.TokenUser{},err
		}
		if referredUserId != 0{
			referralAmount := 150
			err := repository.UpdateReferralAmount(float64(referralAmount),referredUserId,userData.Id)
			if err != nil{
				return &models.TokenUser{},err
			}

			referreason := "Amount credited for used referral code"
			err = repository.UpdateHistory(userData.Id,0,float64(referralAmount),referreason)
			if err != nil{
				return &models.TokenUser{},err
			}

			amount,err := repository.AmountInrefferals(userData.Id)
			if err != nil{
				return &models.TokenUser{},err
			}

		}

	}
	return
}
