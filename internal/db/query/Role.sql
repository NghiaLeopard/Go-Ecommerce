-- name: CreateRole :one
INSERT INTO "Role" (
  name, permission
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetRole :one
SELECT * FROM "Role"
WHERE id = $1 LIMIT 1;

-- name: ListRole :many
SELECT * FROM "Role";

-- name: DeleteRole :exec
DELETE FROM "Role"
WHERE id = $1;