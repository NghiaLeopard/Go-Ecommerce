-- name: CreateUser :one
INSERT INTO "Users" (
  email, password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: InitDefaultAdmin :one
INSERT INTO "Users" (
  email, password, role
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUserById :one
SELECT "Users".*,"Role".*
FROM "Users"
JOIN "Role" ON "Role".id = "Users".role
WHERE "Users".id = $1;

-- name: GetUserByEmail :one
SELECT "Users".*,"Role".*
FROM "Users"
JOIN "Role" ON "Role".id = "Users".role
WHERE "Users".email = $1;

-- name: ListUsers :many
SELECT * FROM "Users"
ORDER BY create_at DESC;

-- name: UpdateUser :one
UPDATE "Users" SET "firstName" = $1,"lastName" = $2,"middleName" = $3, "phoneNumber" = $4,avatar = $5,address = $6,city = $7
WHERE id = $8
RETURNING *;

-- name: UpdatePasswordUser :exec
UPDATE "Users" SET password = $1
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM "Users"
WHERE id = $1;