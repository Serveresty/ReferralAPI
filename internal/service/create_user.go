package service

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"ReferralAPI/pkg/password"
)

func (rs *RefService) CreateUser(user models.User) error {
	var err error

	if user.Password != user.ConfirmPassword {
		return cerr.ErrPasswordsNotEqual
	}

	user.Password, err = password.HashPassword(user.Password)
	if err != nil {
		return err
	}

	err = rs.refRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
