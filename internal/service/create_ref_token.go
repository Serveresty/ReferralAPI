package service

import (
	"strconv"
)

func (rs *RefService) CreateRefToken(id string) (string, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	refCode, err := rs.refRepo.CreateRefToken(idInt)
	if err != nil {
		return "", err
	}

	return refCode, nil
}
