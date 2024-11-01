package service

import "strconv"

func (rs *RefService) DeleteRefToken(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = rs.refRepo.DeleteRefToken(idInt)
	if err != nil {
		return err
	}

	return nil
}
