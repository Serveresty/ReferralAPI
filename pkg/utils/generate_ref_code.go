package utils

import (
	"ReferralAPI/models"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
)

func GenerateRefCode(id int) models.RefCode {
	uid := uuid.New().String()
	ref := base64.RawURLEncoding.EncodeToString([]byte(uid))[:10]
	refCode := models.RefCode{
		ID:          id,
		ReferralStr: ref,
		ExpiredAt:   time.Now().Add(14 * 24 * time.Hour),
	}
	return refCode
}
