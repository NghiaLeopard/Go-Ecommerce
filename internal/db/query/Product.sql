-- name: CreateProduct :one
INSERT INTO "Product" (
  "name","image","countInStock","description","type","status","slug","price","discount","discountStartDate","discountEndDate","location"
) VALUES (
  $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
)
RETURNING *;

-- name: GetProductById :one
SELECT p.*,COUNT(l."user_id") AS "totalLikes",
json_agg(l."user_id") AS "likedBy",
json_agg(v."user_id") AS "uniqueViews" FROM "Product" p
LEFT JOIN "Product_liked" l ON l."product_id" = p.id
LEFT JOIN "Product_UniqueView" v ON v."product_id" = p.id
WHERE p.id = $1 
GROUP BY p.id
LIMIT 1;

-- name: GetProductBySlug :one
SELECT p.*,COUNT(l."user_id") AS "totalLikes",
json_agg(l."user_id") AS "likedBy",
json_agg(v."user_id") AS "uniqueViews" FROM "Product" p
LEFT JOIN "Product_liked" l ON l."product_id" = p.id
LEFT JOIN "Product_UniqueView" v ON v."product_id" = p.id
WHERE p.slug = $1 
GROUP BY p.id
LIMIT 1;

-- name: UpdateViewProduct :one
UPDATE "Product" SET views = $1
WHERE id = $1
RETURNING *;

-- name: DeleteProductById :exec
DELETE FROM "Product"
WHERE id = $1;



-- name: DeleteManyProductsByIds :exec
DELETE FROM "Product"
WHERE id = ANY($1::bigint[]);