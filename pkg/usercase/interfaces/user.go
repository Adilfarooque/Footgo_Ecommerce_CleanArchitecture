package interfaces

import "github.com/Adilfarooque/Footgo_Ecommerce/pkg/utils/models"

type UserUseCase interface {
	UsersSignUP(user models.UserSignUp)(*models.Token,error)
}