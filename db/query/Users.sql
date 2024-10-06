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

-- name: FindUserById :exec
SELECT * FROM "Users"
WHERE id = $1;

-- name: GetUserByEmail :one
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

-- name: UpdateAuthMe :one
UPDATE "Users" SET "avatar" = $1,"address" = $2, "city" = $3,"firstName" = $4,"lastName" = $5,"middleName" = $6,"phoneNumber" = $7,"image" = $8
WHERE id = $9
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "Users"
WHERE id = $1;