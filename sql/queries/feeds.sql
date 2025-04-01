-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    (SELECT id FROM users WHERE users.name = $6)
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * from feeds;

-- name: GetUserByFeedID :one
SELECT * FROM feeds f, users u WHERE f.user_id = u.id AND f.user_id = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeedByID :one
SELECT * FROM feeds WHERE id = $1;