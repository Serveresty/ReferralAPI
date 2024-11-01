package repository

import "context"

func (rr *RefRepository) GetUserIDByRef(refCode string) (int, error) {
	row := rr.db.QueryRow(context.Background(), `SELECT user_id FROM "ref_tokens" WHERE ref_token = $1`, refCode)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
