-- name: NewUser :one
INSERT INTO users (ip)
VALUES ($1)
RETURNING *;
-- name: GetUserByIP :one
SELECT *
FROM users
WHERE ip = sqlc.arg(ip);
-- name: ChangeUsername :one
UPDATE users
SET username = sqlc.arg(username)
WHERE ip = sqlc.arg(ip)
RETURNING *;