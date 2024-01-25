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

func CreateReferralEntry(userDetails models.UserDetailsRespones, userReferral string) error {
	err := db.DB.Exec("INSERT INTO refferrals (user_id,refferal_code,refferal_amount)VALUES(?,?,?)", userDetails.Id, userReferral, 0).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserIdFromReferrals(RefferalCode string) (int, error) {
	var referredUserId int
	err := db.DB.Raw("SELECT user_id FROM referrals WHERE referral_code = ?", RefferalCode).Scan(&referredUserId)
	if err != nil {
		return 0, nil
	}
	return referredUserId, nil
}

func UpdateReferralAmount(referralAmount float64, referredUserId int, currentUserID int) error {
	err := db.DB.Exec("UPDATE referrals SET referral_amount = ? , reffered_user_id = ? ", referralAmount, referredUserId, currentUserID).Error
	if err != nil {
		return err
	}
	err = db.DB.Exec("UPDATE referrals SET referral_amount = referral_amount + ? WHERE user_id = ?", referralAmount, referredUserId).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateHistory(userID, orderID int,amount float64,reason string)error{
	err := db.DB.Exec("INSERT INTO wallet_histories (user_id,orderid,description,amount)VALUES (?,?,?,?)",userID,orderID,reason,amount).Error
	if err != nil{
		return err
	}
	return nil
}

func AmountInrefferals(userID int)(float64,error){
	var a float64
	err := db.DB.Exec("SELECT COUNT(*) FORM wallets WHERE user_id = ?",userID).Scan(&a).Error
	if err != nil{
		return 0.0,err
	}
}