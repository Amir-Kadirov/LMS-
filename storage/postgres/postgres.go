package postgres

import (
	"backend_course/lms/config"
	"backend_course/lms/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

func New(cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return Store{
		DB: db,
	}, nil
}
func (s Store) CloseDB() {
	s.DB.Close()
}

func (s Store) StudentStorage() storage.StudentStorage {
	newStudent := NewStudent(s.DB)

	return &newStudent
}