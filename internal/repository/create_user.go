package repository

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"context"
	"log"
)

func (rr *RefRepository) CreateUser(user models.User) error {
	tx, err := rr.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(context.Background()); rerr != nil {
				log.Printf("rollback error: %v", rerr)
			}
		}
	}()

	if user.ReferralCode != "" {
		refID, err := rr.GetUserIDByRef(user.ReferralCode)
		if err != nil {
			return err
		}

		_, err = tx.Exec(context.Background(), `INSERT INTO "users_data" (email, password, referrer_id) VALUES($1,$2,$3)`, user.Email, user.Password, refID)
		if err != nil {
			return cerr.ErrAlreadyRegistered
		}
	} else {
		_, err = tx.Exec(context.Background(), `INSERT INTO "users_data" (email, password) VALUES($1,$2)`, user.Email, user.Password)
		if err != nil {
			return cerr.ErrAlreadyRegistered
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
