package sqlite

import (
	"BotStudyPSUTI/storage"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

func New(path string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}
	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) Save(order string, user *storage.UserInfo) error {
	return nil
}

func (s *SQLiteStorage) Get(id int) (string, error) {
	return "", nil
}

func (s *SQLiteStorage) Print(query string) (string, error) {
	return "", nil
}
