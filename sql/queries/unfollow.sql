-- name: Unfollow :exec
DELETE from feed_follows where feed_follows.user_id = $1 and feed_follows.feed_id = $2;
