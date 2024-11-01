package models

type UsersList struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	ReferralCode string `json:"referral_code"`
}
