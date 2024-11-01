package models

import "time"

type RefCode struct {
	ID          int       `json:"id"`
	ReferralStr string    `json:"referral_str"`
	ExpiredAt   time.Time `json:"expired_at"`
}
