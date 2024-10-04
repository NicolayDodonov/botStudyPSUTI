package sqlite

import (
	"BotStudyPSUTI/storage"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
)

// SQLiteStorage структура для работы с базой данных.
type SQLiteStorage struct {
	db *sql.DB
}

// New создаёт новый экземпляр SQLiteStorage или nil и ошибку
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

// Save сохраняет запись в order database.db.
func (s *SQLiteStorage) Save(order string, user *storage.UserInfo) error {
	q := `INSERT INTO order (user_order, user_name, type_app) values (?, ?)`
	_, err := s.db.Exec(q, order, user.Username, user.TypeApplication)

	if err != nil {
		return fmt.Errorf("can't insert data into database: %w", err)
	}

	return nil
}

// Print выводит всё содержимое orders из database.db.
func (s *SQLiteStorage) Print(query string) (string, error) {
	type order struct {
		id        int
		userOrder string
		userName  string
		typeApp   string
	}
	orderToString := func(o *[]order) string {
		//Я не уверен в правильности работы этой функции
		var builder strings.Builder
		for i := 0; i < len(*o); i++ {
			builder.WriteString(strconv.Itoa((*o)[i].id))
			builder.WriteString(" ")
			builder.WriteString((*o)[i].userOrder)
			builder.WriteString(" ")
			builder.WriteString((*o)[i].userName)
			builder.WriteString(" ")
			builder.WriteString((*o)[i].typeApp)
			builder.WriteString("\n")
		}
		return builder.String()
	}

	q := `SELECT * FROM order`
	rows, err := s.db.Query(q)
	if err != nil {
		return "", fmt.Errorf("can't get row from database: %w", err)
	}
	defer rows.Close()

	orders := []order{}
	for rows.Next() {
		p := order{}
		err := rows.Scan(&p.id, &p.userOrder, &p.userName, &p.typeApp)
		if err != nil {
			fmt.Println(err)
			continue
		}
		orders = append(orders, p)
	}
	return orderToString(&orders), nil
}

// Init проверяет существование базовой таблицы order в database.db. Если его нет, создаёт таблицу с индексом.
func (s *SQLiteStorage) Init() error {
	q1 := `CREATE TABLE IF NOT EXISTS order (id INT, user_order TEXT, user_name TEXT, type_app TEXT)`
	_, err := s.db.Exec(q1)
	if err != nil {
		return fmt.Errorf("can't init database: ", err)
	}

	q2 := `CREATE INDEX IF NOT EXISTS id ON order(id)`
	_, err = s.db.Exec(q2)
	if err != nil {
		return fmt.Errorf("can't init database: ", err)
	}

	return nil
}
