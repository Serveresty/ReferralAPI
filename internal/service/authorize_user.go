package service

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"ReferralAPI/pkg/password"
	"strconv"
)

func (rs *RefService) AuthorizeUser(user models.User) (string, error) {
	usersData, err := rs.refRepo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if !password.CheckPasswordHash(user.Password, usersData.Password) {
		return "", cerr.ErrWrongPassword
	}

	idStr := strconv.Itoa(usersData.ID)

	token, err := rs.CreateJWTToken(idStr)
	if err != nil {
		return "", err
	}

	return token, nil
}
