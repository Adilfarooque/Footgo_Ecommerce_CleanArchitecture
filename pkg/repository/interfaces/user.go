package interfaces

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/utils/models"
)

type UserRepository interface {
	CheckUserExistsByEmail(email string) (*domain.User, error)
	CheckUserExistsByPhone(phone string) (*domain.User, error)
	UserSignUp(user models.UserSignUp) (models.UserDetailsResponse,error)
	
}
