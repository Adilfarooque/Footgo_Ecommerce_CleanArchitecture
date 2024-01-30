package repository

import (
	"os/user"

	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository (DB *gorm.DB)interfaces.UserRepository{
	return &userRepository{
		DB:DB,
	}
}

func (ur *userRepository) CheckUserExistsByEmail(email string)(*domain.user,error){
	
}