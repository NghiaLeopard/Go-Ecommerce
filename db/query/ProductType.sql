-- name: CreateProductType :one
INSERT INTO "Product_Type" (
  name,slug
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetProductTypeById :one
SELECT * FROM "Product_Type"
WHERE id = $1 LIMIT 1;

-- name: GetProductTypeByName :one
SELECT * FROM "Product_Type"
WHERE name = $1 LIMIT 1;

-- name: ListProductType :many
SELECT *,COUNT("Product_Type".id) OVER() AS "totalCount" FROM "Product_Type"
WHERE  @search ::text = '' or name ILIKE concat('%',@search,'%')
ORDER BY 
  CASE 
        WHEN @order_by ::varchar = 'name asc' THEN name END ASC,
  CASE 
        WHEN @order_by = 'name desc' THEN name END DESC,
  CASE 
        WHEN @order_by = 'slug asc' THEN slug END ASC,
  CASE 
        WHEN @order_by = 'slug desc' THEN slug END DESC,
  CASE 
        WHEN @order_by = 'created_date asc' THEN create_at END ASC,
  CASE 
        WHEN @order_by = 'created_date desc' THEN create_at END DESC
LIMIT $1
OFFSET $2;

-- name: UpdateProductType :one
UPDATE "Product_Type" SET name = $1,slug = $2,update_at = NOW()
WHERE id = $3
RETURNING *;

-- name: DeleteProductTypeById :exec
DELETE FROM "Product_Type"
WHERE id = $1;

-- name: DeleteManyProductTypesByIds :exec
DELETE FROM "Product_Type"
WHERE id = ANY($1::bigint[]);

