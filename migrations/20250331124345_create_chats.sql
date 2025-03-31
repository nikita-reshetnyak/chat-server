-- +goose Up
CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT now()
);
-- +goose Down
DROP TABLE chats;