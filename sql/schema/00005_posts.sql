-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
    id           UUID PRIMARY KEY,
    feed_id      UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    url          TEXT UNIQUE NOT NULL,
    title        TEXT,
    description  TEXT,
    published_at TIMESTAMP WITH TIME ZONE,
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE posts;
