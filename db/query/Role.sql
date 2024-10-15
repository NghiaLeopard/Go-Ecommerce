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
SELECT _id,name,coalesce(permission, ARRAY[]::TEXT[]) FROM "Role"
WHERE _id = $1 LIMIT 1;

-- name: GetRoleByName :one
SELECT * FROM "Role"
WHERE name = $1 LIMIT 1;

-- name: ListRole :many
SELECT *,COUNT("Role"."_id") OVER() AS "totalCount" FROM "Role"
WHERE  @search ::text = '' or name ILIKE concat('%',@search,'%')
LIMIT NULLIF(@limit_opt :: int, 0)
OFFSET NULLIF(@offset_opt :: int, 0);

-- name: UpdateRole :one
UPDATE "Role" SET name = $1,permission = $2,update_at = NOW()
WHERE "_id" = $3
RETURNING *;

-- name: DeleteRoleById :exec
DELETE FROM "Role"
WHERE "_id" = $1;

-- name: DeleteManyRolesByIds :exec
DELETE FROM "Role"
WHERE "_id" = ANY($1::bigint[]);

