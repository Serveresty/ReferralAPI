package repository

import (
	cerr "ReferralAPI/pkg/custom_errors"
	"context"
)

func (rr *RefRepository) GetTokenByEmail(email string) (string, error) {
	row := rr.db.QueryRow(context.Background(), `SELECT rt.ref_token FROM ref_tokens rt JOIN users_data u ON rt.user_id = u.user_id WHERE u.email = $1;`, email)

	var token string
	err := row.Scan(&token)
	if err != nil {
		return "", cerr.ErrReferralTokenNotFound
	}

	return token, nil
}
