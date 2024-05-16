package postgres

import (
	"backend_course/lms/config"
	"backend_course/lms/storage"
	"backend_course/lms/storage/redis"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

type Store struct {
	Pool  *pgxpool.Pool
	cfg config.Config
	redis storage.IRedisStorage
}

func New(ctx context.Context, cfg config.Config, redis storage.IRedisStorage) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	pgxPoolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	pgxPoolConfig.MaxConns = 50
	pgxPoolConfig.MaxConnLifetime = time.Hour

	newPool, err := pgxpool.NewWithConfig(ctx, pgxPoolConfig)
	if err != nil {
		return nil, err
	}

	return Store{
		Pool:  newPool,
		redis: redis,
		cfg:   cfg,
	}, nil
}
func (s Store) CloseDB() {
	s.Pool.Close()
}

func (s Store) StudentStorage() storage.StudentStorage {
	newStudent := NewStudent(s.Pool)

	return &newStudent
}

func (s Store) TeacherStorage() storage.TeacherStorage {
	newTeacher := NewTeacher(s.Pool)

	return &newTeacher
}

func (s Store) SubjectStorage() storage.SubjectStorage {
	newSubject := NewSubject(s.Pool)

	return &newSubject
}

func (s Store) TimeTableStorage() storage.TimeTableStorage {
	newTimeTable := NewTimeTable(s.Pool)

	return &newTimeTable
}

func (s Store) Redis() storage.IRedisStorage {
	return redis.New(s.cfg)
}
