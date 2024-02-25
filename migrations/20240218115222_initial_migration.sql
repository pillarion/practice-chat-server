-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    from_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS chats_users (
    chat_id INTEGER REFERENCES chats(id) NOT NULL,
    user_id INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS chats_users;
DROP TABLE IF EXISTS chats;
-- +goose StatementEnd