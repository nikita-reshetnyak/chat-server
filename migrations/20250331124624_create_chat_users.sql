-- +goose Up
CREATE TABLE chat_users (
    chat_id INT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (chat_id, user_id)
);
-- +goose Down
DROP TABLE chat_users;
