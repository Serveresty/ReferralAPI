package models

type User struct {
	ID              int    `json:"id"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	ReferralCode    string `json:"referral_code"`
}
