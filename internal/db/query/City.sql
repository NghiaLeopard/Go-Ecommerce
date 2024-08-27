-- name: CreateCity :one
INSERT INTO "City" (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetCity :one
SELECT * FROM "City"
WHERE id = $1 LIMIT 1;

-- name: ListCity :many
SELECT * FROM "City"
WHERE name ILIKE '%' || $1 || '%'
ORDER BY $2
LIMIT $3
OFFSET $4;

-- name: UpdateCity :one
UPDATE "City" SET name = $1
WHERE id = $2
RETURNING *;

-- name: DeleteCity :exec
DELETE FROM "City"
WHERE id = $1;