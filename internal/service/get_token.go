package service

func (rc *RefService) GetToken(email string) (string, error) {
	token, err := rc.refRepo.GetTokenByEmail(email)
	if err != nil {
		return "", err
	}

	return token, nil
}
