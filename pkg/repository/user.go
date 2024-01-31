package repository

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/repository/interfaces"
	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (ur *userRepository) CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := ur.DB.Where(&domain.User{Email: email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func (ur *userRepository) CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User
	res := ur.DB.Where(&domain.User{Phone: phone}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func (ur *userRepository) UserSignUp(user models.UserSignUp) (models.UserDetailsResponse, error) {
	var SignupDetails models.UserDetailsResponse
	err := ur.DB.Raw("INSERT INTO users(firstname,lastname,email,password,phone)VALUES(?,?,?,?,?)RETURNING id,firstname,lastname,email,password,phone", user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).Scan(&SignupDetails).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	return SignupDetails, nil

}

