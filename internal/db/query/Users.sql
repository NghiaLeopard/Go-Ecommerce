-- name: CreateUser :one
INSERT INTO "Users" (
  email, password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: InitDefaultAdmin :one
INSERT INTO "Users" (
  email, password,role
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "Users"
WHERE id = $1 LIMIT 1;

-- name: FindEmail :one
SELECT "Users".*,"Role".*
FROM "Users"
JOIN "Role" ON "Role".id = "Users".role
WHERE "Users".email = $1;

-- name: ListUsers :many
SELECT * FROM "Users"
ORDER BY create_at DESC;

-- name: UpdatePasswordUser :exec
UPDATE "Users" SET password = $1
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM "Users"
WHERE id = $1;