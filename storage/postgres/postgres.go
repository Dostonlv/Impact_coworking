package postgres

import (
	"Impact/config"
	"Impact/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db      *pgxpool.Pool
	room    storage.RoomRepoI
	booking storage.BookingRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Room() storage.RoomRepoI {
	if s.room == nil {
		s.room = NewRoomRepo(s.db)
	}
	return s.room
}

func (s *Store) Booking() storage.BookingRepoI {
	if s.booking == nil {
		s.booking = NewBookingRoomRepo(s.db)
	}
	return s.booking
}
