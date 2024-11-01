package repository

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"context"
)

func (rr *RefRepository) GetUserByEmail(email string) (models.User, error) {
	row := rr.db.QueryRow(context.Background(), `SELECT user_id, password FROM "users_data" WHERE email = $1`, email)

	var user models.User
	user.Email = email
	err := row.Scan(&user.ID, &user.Password)
	if err != nil {
		return models.User{}, cerr.ErrUserNotFound
	}

	return user, nil
}
