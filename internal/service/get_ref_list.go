package service

import (
	"ReferralAPI/models"
	"strconv"
)

func (rs *RefService) GetRefList(id string) ([]models.UsersList, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return []models.UsersList{}, err
	}

	users, err := rs.refRepo.GetRefList(idInt)
	if err != nil {
		return users, err
	}

	return users, nil
}
