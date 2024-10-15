-- name: CreateDelivery :one
INSERT INTO "Delivery_Type" (
  name,price
) VALUES (
  $1,$2
)
RETURNING *;

-- name: GetDeliveryById :one
SELECT * FROM "Delivery_Type"
WHERE "_id" = $1 LIMIT 1;

-- name: GetDeliveryByName :one
SELECT * FROM "Delivery_Type"
WHERE name = sqlc.arg(name) LIMIT 1;

-- name: ListDelivery :many
SELECT *,COUNT("Delivery_Type"."_id") OVER() AS "totalCount" FROM "Delivery_Type"
WHERE  @search ::text = '' or name ILIKE concat('%',@search,'%')
ORDER BY 
  CASE 
        WHEN @order_by ::varchar = 'name asc' THEN name END ASC,
  CASE 
        WHEN @order_by = 'name desc' THEN name END DESC
LIMIT NULLIF(@limit_opt :: int, 0)
OFFSET NULLIF(@offset_opt :: int, 0);


-- name: UpdateDelivery :one
UPDATE "Delivery_Type" SET name = $1,price = $2,update_at = NOW()
WHERE "_id" = $3
RETURNING *;

-- name: DeleteDeliveryById :exec
DELETE FROM "Delivery_Type"
WHERE "_id" = $1;

-- name: DeleteManyDeliveryByIds :exec
DELETE FROM "Delivery_Type"
WHERE "_id" = ANY($1::bigint[]);