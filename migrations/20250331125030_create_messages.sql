-- +goose Up
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    chat_id INT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    sender_id INT NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    text TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE messages;
