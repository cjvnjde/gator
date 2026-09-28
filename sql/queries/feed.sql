-- name: CreateFeed :one
INSERT INTO
  feeds (id, created_at, updated_at, url, user_id, name)
VALUES
  ($1, $2, $3, $4, $5, $6)
RETURNING
  *;
