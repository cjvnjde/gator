-- name: GetFeedFollowForUser :many
SELECT
  feed_follows.*,
  sqlc.embed(users),
  sqlc.embed(feeds)
FROM
  feed_follows
  INNER JOIN users ON users.id = feed_follows.user_id
  INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE
  feed_follows.user_id = $1;
