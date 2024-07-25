package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type MessageRepositoryInterface interface {
	Create(message string) (int, error)
	GetAll() (map[string]int, error)
}

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (r MessageRepository) Create(message string) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO messages(message, status) VALUES($1, 'waiting') RETURNING id", message).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r MessageRepository) GetAll() (map[string]int, error) {
	var counts = map[string]int{
		"waiting": 0,
		"success": 0,
		"failed":  0,
	}

	rows, err := r.db.Query("SELECT status, COUNT(*) FROM messages GROUP BY status")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var status string
		var count int
		err := rows.Scan(&status, &count)
		if err != nil {
			return nil, err
		}
		counts[status] = count
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return counts, nil
}
