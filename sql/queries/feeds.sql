-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: FetchFeedsByUser :many
select id, created_at, updated_at, name, url, user_id from feeds where user_id = $1;
