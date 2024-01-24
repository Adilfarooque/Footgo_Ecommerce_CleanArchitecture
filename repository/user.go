package repository

import (
	"errors"

	"github.com/Adilfarooque/Footsgo_Ecommerce/db"
	"github.com/Adilfarooque/Footsgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footsgo_Ecommerce/utils/models"
	"gorm.io/gorm"
)

func CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := db.DB.Where(domain.User{Email: email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil

}

func CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User
	res := db.DB.Where(&domain.User{Phone: phone}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func UserSignUp(user models.UserSignUp) (models.UserDetailsRespones, error) {
	var SignupDetail models.UserDetailsRespones
	err := db.DB.Raw("INSERT INTO users(firstname,lastname,email,password,phone)VALUE(?,?,?,?,?)RETURNING id,firstname,lastname,email,password,phone", user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).Scan(&SignupDetail).Error
	if err != nil {
		return models.UserDetailsRespones{}, err
	}
	return SignupDetail, nil

}
