-- name: CreatePost :exec
INSERT INTO posts (
    id,
    feed_id,
    url,
    title,
    description,
    published_at,
    created_at,
    updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, DEFAULT, DEFAULT)
ON CONFLICT (url) DO NOTHING;

-- name: GetPostsForUser :many
SELECT p.*
FROM posts p
INNER JOIN feeds f ON f.id = p.feed_id
INNER JOIN users u ON u.id = f.user_id
WHERE u.name = $1
ORDER BY f.last_fetched_at NULLS FIRST
LIMIT $2;
