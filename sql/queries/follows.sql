-- name: CreateFollow :one
INSERT INTO feed_follows (id, feed_id, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: FetchFollowsByUser :many
select id,  feed_id, user_id from feed_follows where user_id = $1;

-- name: FetchFeedsFollowedByUser :many
select id, created_at, updated_at, name, url, user_id from feeds where id in (select   feed_id  from feed_follows ff where ff.user_id = $1);
