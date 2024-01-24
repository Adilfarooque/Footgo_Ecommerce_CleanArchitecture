package models

type UserSignUp struct {
	Firstname    string `json:"firstname validate:"gte=3"`
	Lastname     string `json:"lastname" validate:"gte=1"`
	Email        string `json:"email" validate:"email"`
	Password     string `json:"password" validate:"min=6,max20"`
	Phone        string `json:"phone" validate:"e164"`
	ReferralCode string `json:"referral_code`
}

type UserDetailsRespones struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserDetailsAtAdmin struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Blocked   bool   `json:"blocked"`
}

type TokenUser struct {
	Users        UserDetailsRespones
	AccessToken  string
	RefreshToken string
}
