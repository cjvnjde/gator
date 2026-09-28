-- name: GetFeed :one
SELECT
  feeds.*,
  sqlc.embed(users)
FROM
  feeds
  JOIN users ON feeds.user_id = users.id
WHERE
  feeds.url = $1;
