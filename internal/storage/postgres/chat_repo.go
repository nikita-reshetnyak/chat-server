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
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, err //handle
	}
	defer tx.Rollback(ctx)
	//todo add user if not exist
	userIDs := make([]int64, 0, len(usernames))
	for _, username := range usernames {
		var userID int64
		err := tx.QueryRow(ctx, `
		INSERT INTO users (username)
		VALUES ($1)
		ON CONFLICT (username) DO UPDATE SET username = EXCLUDED.username
		RETURNING id`,
			username).Scan(&userID)
		if err != nil {
			return 0, fmt.Errorf("insert user: %d: %w", userID, err)
		}
		userIDs = append(userIDs, userID)
	}
	//todo create chat
	var chatId int64
	err = tx.QueryRow(ctx, `
	INSERT INTO chats DEFAULT VALUES RETURNING id
	`).Scan(&chatId)
	if err != nil {
		return 0, fmt.Errorf("create chat: %w", err)
	}
	//todo connect users to chat_users
	for _, userId := range userIDs {
		_, err := tx.Exec(ctx, "INSERT INTO chat_users (chat_id,user_id) VALUES ($1,$2)", chatId, userId)
		if err != nil {
			return 0, fmt.Errorf("link user:%d into chat:%d: %w", userId, chatId, err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("commit: %w", err)
	}

	return chatId, nil
}

func (s *Storage) Delete(ctx context.Context, id int64) error {
	// твоя логика
	return nil
}

func (s *Storage) SendMessage(ctx context.Context, from string, text string) error {
	// твоя логика
	return nil
}
