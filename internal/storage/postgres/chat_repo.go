package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

func New(connString string) (*Storage, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &Storage{db: conn}, nil
}
func (s *Storage) Create(ctx context.Context, usernames []string) (int64, error) {
	return 1, nil
}

func (s *Storage) Delete(ctx context.Context, id int64) error {
	// твоя логика
	return nil
}

func (s *Storage) SendMessage(ctx context.Context, from string, text string) error {
	// твоя логика
	return nil
}
