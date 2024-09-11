-- name: CreateRole :one
INSERT INTO "Role" (
  name, permission
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetRoleById :one
SELECT * FROM "Role"
WHERE id = $1 LIMIT 1;

-- name: GetRoleByName :one
SELECT * FROM "Role"
WHERE name = $1 LIMIT 1;

-- name: ListRole :many
SELECT * FROM "Role";

-- name: UpdateRole :one
UPDATE "Role" SET name = $1,permission = $2
WHERE id = $3
RETURNING *;

-- name: DeleteRoleById :exec
DELETE FROM "Role"
WHERE id = $1;

-- name: DeleteManyRolesByIds :exec
DELETE FROM "Role"
WHERE id = ANY($1::bigint[]);

