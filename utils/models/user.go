package models

type UserSignUp struct {
	Firstname    string `json:"firstname validate:"gte=3"`
	Lastname     string	`json:"lastname" validate:"gte=1"`
	Email        string	`json:"email" validate:"email"`
	Password     string	`json:"password" validate:"min=6,max20"`
	Phone        string	`json:"phone" validate:"e164"`
	ReferralCode string	`json:"referral_code`
}
