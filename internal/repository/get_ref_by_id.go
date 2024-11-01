package repository

import "context"

func (rr *RefRepository) GetRefCodeById(id int) (string, error) {
	row := rr.db.QueryRow(context.Background(), `SELECT ref_token FROM "ref_tokens" WHERE user_id = $1`, id)
	var code string
	err := row.Scan(&code)
	if err != nil {
		return "", err
	}

	return code, nil
}
