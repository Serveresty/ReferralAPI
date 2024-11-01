package repository

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"ReferralAPI/pkg/utils"
	"context"
	"database/sql"
	"errors"
	"log"
)

func (rr *RefRepository) CreateRefToken(id int) (string, error) {
	tx, err := rr.db.Begin(context.Background())
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(context.Background()); rerr != nil {
				log.Printf("rollback error: %v", rerr)
			}
		}
	}()

	var ref models.RefCode
	var retrier bool
	for i := 0; i < 5; i++ {
		ref = utils.GenerateRefCode(id)
		_, eerr := rr.GetUserIDByRef(ref.ReferralStr)
		if eerr != nil {
			if errors.Is(eerr, sql.ErrNoRows) {
				retrier = false
				break
			}
			return "", eerr
		}
		retrier = true
		continue
	}
	if retrier {
		return "", cerr.ErrTooManyRetries
	}

	_, err = rr.GetRefCodeById(ref.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	if errors.Is(err, sql.ErrNoRows) {
		_, err = tx.Exec(context.Background(), `INSERT INTO "ref_tokens" (ref_token, user_id, expired_at) VALUES($1,$2,$3)`, ref.ReferralStr, ref.ID, ref.ExpiredAt)
		if err != nil {
			return "", err
		}
	} else {
		_, err = tx.Exec(context.Background(), `UPDATE "ref_tokens" SET ref_token = $1, expired_at = $2 WHERE user_id = $3`, ref.ReferralStr, ref.ExpiredAt, ref.ID)
		if err != nil {
			return "", err
		}
	}

	errr := tx.Commit(context.Background())
	if errr != nil {
		return "", errr
	}

	return ref.ReferralStr, nil
}
