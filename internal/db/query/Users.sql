-- name: CreateUser :one
INSERT INTO "Users" (
  email, password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "Users"
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "Users"
ORDER BY create_at DESC;

-- name: DeleteUser :exec
DELETE FROM "Users"
WHERE id = $1;