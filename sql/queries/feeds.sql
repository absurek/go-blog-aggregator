-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, name, url, created_at, updated_at)
VALUES ($1, $2, $3, $4, DEFAULT, DEFAULT)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name, f.url, u.name AS user_name
FROM users u
INNER JOIN feeds f ON f.user_id = u.id;

-- name: GetFeedByURL :one
SELECT id, user_id, name, url, created_at, updated_at
FROM feeds
WHERE url = $1;
