-- name: CreatePayment :one
INSERT INTO "Payment_Type" (
  name,type
) VALUES (
  $1,$2
)
RETURNING *;

-- name: GetPaymentById :one
SELECT * FROM "Payment_Type"
WHERE "_id" = $1 LIMIT 1;

-- name: GetPaymentByName :one
SELECT * FROM "Payment_Type"
WHERE name = sqlc.arg(name) LIMIT 1;

-- name: ListPayment :many
SELECT *,COUNT("Payment_Type"."_id") OVER() AS "totalCount" FROM "Payment_Type"
WHERE  @search ::text = '' or name ILIKE concat('%',@search,'%')
ORDER BY 
  CASE 
        WHEN @order_by ::varchar = 'name asc' THEN name END ASC,
  CASE 
        WHEN @order_by = 'name desc' THEN name END DESC
LIMIT NULLIF(@limit_opt :: int, 0)
OFFSET NULLIF(@offset_opt :: int, 0);


-- name: UpdatePayment :one
UPDATE "Payment_Type" SET name = $1,type = $2,update_at = NOW()
WHERE "_id" = $3
RETURNING *;

-- name: DeletePaymentById :exec
DELETE FROM "Payment_Type"
WHERE "_id" = $1;

-- name: DeleteManyPaymentByIds :exec
DELETE FROM "Payment_Type"
WHERE "_id" = ANY($1::bigint[]);