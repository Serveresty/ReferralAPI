package repository

import (
	"context"
	"log"
)

func (rr *RefRepository) DeleteRefToken(id int) error {
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

	_, err = tx.Exec(context.Background(), `DELETE FROM "ref_tokens" WHERE user_id = $1`, id)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
