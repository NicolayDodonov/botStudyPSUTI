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
	q := `INSERT INTO order (user_order, user_id) values (?, ?)`
	_, err := s.db.Exec(q, order, user.UserID)

	if err != nil {
		return fmt.Errorf("can't insert data into database: %w", err)
	}

	return nil
}

func (s *SQLiteStorage) Get(id int) (string, error) {
	return "", nil
}

func (s *SQLiteStorage) Print(query string) (string, error) {
	return "", nil
}

func (s *SQLiteStorage) Init() error {
	q1 := `CREATE TABLE IF NOT EXISTS order (id INT, user_order TEXT, user_id INT)`
	_, err := s.db.Exec(q1)
	if err != nil {
		return fmt.Errorf("can't init database: ", err)
	}

	q2 := `CREATE INDEX IF NOT EXISTS id ON order(id, user_order)`
	_, err = s.db.Exec(q2)
	if err != nil {
		return fmt.Errorf("can't init database: ", err)
	}

	return nil
}
