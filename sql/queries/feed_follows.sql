-- name: CreateFeedFollow :one

WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
    RETURNING *
)

SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsByUser :many
SELECT ff.* FROM feed_follows ff, users u
WHERE u.name = $1
AND ff.user_id = u.id;

-- name: DeleteFeedFollowByUrlForUser :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;