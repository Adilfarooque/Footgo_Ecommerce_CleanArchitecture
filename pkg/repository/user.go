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

func (ur *userRepository) UserSignUp(user models.UserSignUp)(models.UserDetailsResponse,error){
	
}