package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hramcovdv/cm_info/models"
)

type SqliteStorage struct {
	db *sql.DB
}

func NewSqliteStorage(source string) (*SqliteStorage, error) {
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}

	return &SqliteStorage{db: db}, nil
}

func (s *SqliteStorage) GetHeadend(h *models.Headend, addr string) error {
	row := s.db.QueryRow("SELECT address, community FROM headends WHERE address=?", addr)
	if err := row.Scan(&h.Addr, &h.Comm); err != nil {
		return err
	}

	return nil
}

func (s *SqliteStorage) GetHeadends() ([]models.Headend, error) {
	rows, err := s.db.Query("SELECT address, community FROM headends")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var a []models.Headend

	for rows.Next() {
		h := models.Headend{}
		err = rows.Scan(&h.Addr, &h.Comm)
		if err != nil {
			return nil, err
		}

		a = append(a, h)
	}

	return a, nil
}

func (s *SqliteStorage) Close() error {
	return s.db.Close()
}
