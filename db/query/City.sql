-- name: CreateCity :one
INSERT INTO "City" (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetCityById :one
SELECT * FROM "City"
WHERE "_id" = $1 LIMIT 1;

-- name: GetCityByName :one
SELECT * FROM "City"
WHERE name = sqlc.arg(name) LIMIT 1;

-- name: ListCity :many
SELECT *,COUNT("City"."_id") OVER() AS "totalCount" FROM "City"
WHERE  @search ::text = '' or name ILIKE concat('%',@search,'%')
ORDER BY 
  CASE 
        WHEN @order_by ::varchar = 'name asc' THEN name END ASC,
  CASE 
        WHEN @order_by = 'name desc' THEN name END DESC
LIMIT NULLIF(@limit_opt :: int, 0)
OFFSET NULLIF(@offset_opt :: int, 0);


-- name: UpdateCity :one
UPDATE "City" SET name = $1,update_at = NOW()
WHERE "_id" = $2
RETURNING *;

-- name: DeleteCityById :exec
DELETE FROM "City"
WHERE "_id" = $1;

-- name: DeleteManyCityByIds :exec
DELETE FROM "City"
WHERE "_id" = ANY($1::bigint[]);