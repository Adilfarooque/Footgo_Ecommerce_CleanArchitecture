package repository

import (
	"github.com/Adilfarooque/Footsgo_Ecommerce/domain"
)

func CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := db.DB.Where()
}
