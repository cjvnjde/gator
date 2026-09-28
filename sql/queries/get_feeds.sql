-- name: GetFeeds :many

SELECT
  feeds.*,
  sqlc.embed(users)
FROM
  feeds
  JOIN users ON feeds.user_id = users.id;
