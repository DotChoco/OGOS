package animestb

import "database/sql"

func UpdateAnime(db *sql.DB, name string, email string) (int64, error) {
	result, err := db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", name, email)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}
