package Database

import (
	"database/sql"
)

const DBurl = "root:password@tcp(127.0.0.1:3306)/schema_name"

type Store struct {
	db *sql.DB
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", DBurl)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}
func (s *Store) Close() {
	s.db.Close()
}
