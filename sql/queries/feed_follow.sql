-- name: CreateFeedFollow :one
WITH
  insterted_feed_follow AS (
    INSERT INTO
      feed_follows (id, user_id, feed_id, created_at, updated_at)
    VALUES
      ($1, $2, $3, $4, $5)
    RETURNING
      *
  )
SELECT
  insterted_feed_follow.*,
  sqlc.embed(users),
  sqlc.embed(feeds)
FROM
  insterted_feed_follow
  INNER JOIN users ON users.id = insterted_feed_follow.user_id
  INNER JOIN feeds ON feeds.id = insterted_feed_follow.feed_id;
