-- name: CreateCity :one
INSERT INTO "City" (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetCityById :one
SELECT * FROM "City"
WHERE id = $1 LIMIT 1;

-- name: GetCityByName :one
SELECT * FROM "City"
WHERE name = sqlc.arg(name) LIMIT 1;

-- name: ListCity :many
SELECT * FROM "City"
WHERE name ILIKE '%' || $1 || '%'
ORDER BY $2
LIMIT $3
OFFSET $4;

-- name: UpdateCity :one
UPDATE "City" SET name = $1,update_at = NOW()
WHERE id = $2
RETURNING *;

-- name: DeleteCityById :exec
DELETE FROM "City"
WHERE id = $1;

-- name: DeleteManyCityByIds :exec
DELETE FROM "City"
WHERE id = ANY($1::bigint[]);