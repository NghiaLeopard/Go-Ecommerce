-- name: CreateRoleByDefault :one
INSERT INTO "Role" (
  name, permission
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreateRole :one
INSERT INTO "Role" (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetRoleById :one
SELECT * FROM "Role"
WHERE id = $1 LIMIT 1;

-- name: GetRoleByName :one
SELECT * FROM "Role"
WHERE name = $1 LIMIT 1;

-- name: ListRole :many
SELECT * FROM "Role"
WHERE  @search ::text = '' or name ILIKE concat('%',@search,'%')
LIMIT $1
OFFSET $2;

-- name: UpdateRole :one
UPDATE "Role" SET name = $1,permission = $2,update_at = NOW()
WHERE id = $3
RETURNING *;

-- name: DeleteRoleById :exec
DELETE FROM "Role"
WHERE id = $1;

-- name: DeleteManyRolesByIds :exec
DELETE FROM "Role"
WHERE id = ANY($1::bigint[]);

