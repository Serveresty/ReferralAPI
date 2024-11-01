package repository

import (
	"ReferralAPI/models"
	"context"
	"database/sql"
	"log"
)

func (rr *RefRepository) GetRefList(id int) ([]models.UsersList, error) {
	tx, err := rr.db.Begin(context.Background())
	if err != nil {
		return []models.UsersList{}, err
	}
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(context.Background()); rerr != nil {
				log.Printf("rollback error: %v", rerr)
			}
		}
	}()

	rows, err := tx.Query(context.Background(), `SELECT u.user_id, u.email, rt.ref_token FROM users_data u LEFT JOIN ref_tokens rt ON u.user_id = rt.user_id WHERE u.referrer_id = $1;`, id)
	if err != nil {
		return []models.UsersList{}, err
	}

	var users []models.UsersList
	for rows.Next() {
		var user models.UsersList
		var nlstr sql.NullString
		err = rows.Scan(&user.ID, &user.Email, &nlstr)
		if err != nil {
			return []models.UsersList{}, err
		}
		user.ReferralCode = nlstr.String
		users = append(users, user)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return []models.UsersList{}, err
	}

	return users, nil
}
