package animestb

import "database/sql"

func createTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(50),
			email VARCHAR(50),
			PRIMARY KEY (id)
		);`
	_, err := db.Exec(query)
	return err
}

func AddAnime(db *sql.DB, name string, email string) (int64, error) {
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
