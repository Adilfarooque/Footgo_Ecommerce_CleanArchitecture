package models

type UserSignUp struct {
	Firstname    string `json:"firstname" validate:"gte=3"`
	Lastname     string `json:"lastname" validate:"gte=1"`
	Email        string `json:"email" validate:"email"`
	Password     string `json:"password" validate:"min=6",max="20"`
	Phone        string `json:"phone" vaildate:"e164`
	ReferralCode string `json:"referral_code"`
}

type UserDetailsResponse struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserDetailsAtAdmin struct {
	Id          int    `json:"id"`
	Firstname   string `json:"firstname`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	BlockStatus string `json:"block_status`
}

type Token struct {
	Users        UserDetailsResponse
	AccessToken  string
	RefreshToken string
}


