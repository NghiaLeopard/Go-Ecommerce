-- name: CreateProductView :one
INSERT INTO "Product_UniqueView" (
  "product_id","user_id"
) VALUES (
  $1, $2
)
RETURNING *;