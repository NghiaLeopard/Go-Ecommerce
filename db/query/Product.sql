-- name: CreateProduct :one
INSERT INTO "Product" (
  "name","image","countInStock","description","type","status","slug","price","discount","discountStartDate","discountEndDate","location"
) VALUES (
  $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
)
RETURNING *;

-- name: CreateProductUniqueView :exec
INSERT INTO "Product_UniqueView" (
  "product_id","user_id"
) VALUES (
  $1, $2
);

-- name: CreateProductLike :exec
INSERT INTO "Product_liked" (
  "product_id","user_id"
) VALUES (
  $1, $2
);

-- name: UpdateProduct :one
UPDATE "Product" SET name = $1,image = $2,"countInStock" = $3,description = $4,type = $5,status = $6,slug = $7,price = $8,discount = $9,"discountStartDate" = $10,"discountEndDate" = $11,location = $12
WHERE "_id" = $13
RETURNING *;

-- name: GetProductTypeBySlug :one
SELECT "_id",type from "Product"
WHERE slug = $1 LIMIT 1;

-- name: CheckProduct :one
SELECT "_id" FROM "Product"
WHERE "_id" = $1 LIMIT 1;

-- name: GetProductById :one
SELECT p.*,COUNT(l."user_id") AS "totalLikes",
json_agg(l."user_id") AS "likedBy",
json_agg(v."user_id") AS "uniqueViews" FROM "Product" p
LEFT JOIN "Product_liked" l ON l."product_id" = p."_id"
LEFT JOIN "Product_UniqueView" v ON v."product_id" = p."_id"
WHERE p."_id" = $1 
GROUP BY p."_id"
LIMIT 1;

-- name: GetAllProductRelated :many
SELECT p.*,COUNT(l."user_id") AS "totalLikes",COUNT(p."_id") OVER() AS "totalCount",
CASE WHEN COUNT(l."user_id") > 0 THEN json_agg(l."user_id") ELSE '[]'::json END AS "likedBy",
CASE WHEN COUNT(v."user_id") > 0 THEN json_agg(v."user_id") ELSE '[]'::json END AS "uniqueViews"
FROM "Product" p
LEFT JOIN "Product_liked" l ON l."product_id" = p."_id"
LEFT JOIN "Product_UniqueView" v ON v."product_id" = p."_id"
WHERE p.type = $1 AND p."_id" <> $2
GROUP BY p."_id"
LIMIT $3
OFFSET $4;

-- name: GetProductBySlug :one
SELECT p.*,COUNT(l."user_id") AS "totalLikes",
json_agg(l."user_id") AS "likedBy",
json_agg(v."user_id") AS "uniqueViews" FROM "Product" p
LEFT JOIN "Product_liked" l ON l."product_id" = p."_id"
LEFT JOIN "Product_UniqueView" v ON v."product_id" = p."_id"
WHERE p.slug = $1 
GROUP BY p."_id"
LIMIT 1;

-- name: GetProductPublicById :one
SELECT p.*,COUNT(l."user_id") AS "totalLikes",
json_agg(l."user_id") AS "likedBy",
json_agg(v."user_id") AS "uniqueViews" FROM "Product" p
LEFT JOIN "Product_liked" l ON l."product_id" = p."_id"
LEFT JOIN "Product_UniqueView" v ON v."product_id" = p."_id"
WHERE p."_id" = $1 
GROUP BY p."_id"
LIMIT 1;

-- name: GetAllProductLike :many
SELECT p.*,COUNT(l."user_id") AS "totalLikes",COUNT(p."_id") OVER() AS "totalCount",
CASE WHEN COUNT(l."user_id") > 0 THEN json_agg(l."user_id") ELSE '[]'::json END AS "likedBy",
CASE WHEN COUNT(v."user_id") > 0 THEN json_agg(v."user_id") ELSE '[]'::json END AS "uniqueViews"
FROM "Product" p
LEFT JOIN "Product_liked" l ON l."product_id" = p."_id"
LEFT JOIN "Product_UniqueView" v ON v."product_id" = p."_id"
WHERE l.user_id = $1 AND (@search ::text = '' or name ILIKE concat('%',@search,'%'))   
GROUP BY p."_id"
ORDER BY MAX(l.like_date) asc
LIMIT $2
OFFSET $3;

-- name: GetAllProductView :many
SELECT p.*,COUNT(l."user_id") AS "totalLikes",COUNT(p."_id") OVER() AS "totalCount",
CASE WHEN COUNT(l."user_id") > 0 THEN json_agg(l."user_id") ELSE '[]'::json END AS "likedBy",
CASE WHEN COUNT(v."user_id") > 0 THEN json_agg(v."user_id") ELSE '[]'::json END AS "uniqueViews"
FROM "Product" p
LEFT JOIN "Product_liked" l ON l."product_id" = p."_id"
LEFT JOIN "Product_UniqueView" v ON v."product_id" = p."_id"
WHERE v.user_id = $1 AND (@search ::text = '' or name ILIKE concat('%',@search,'%'))
GROUP BY p."_id" 
ORDER BY MAX(v.view_date) asc
LIMIT $2
OFFSET $3;

-- name: GetAllProductAdmin :many
SELECT p."_id",p.name,p."countInStock",p.image,p.price,p.slug,p.status,
json_build_object('_id', pt."_id", 'name', pt.name) AS "type",
COUNT(p."_id") OVER() AS "totalCount" FROM "Product" p
JOIN "Product_Type" pt ON p.type = pt."_id"
WHERE 
  CASE
		WHEN @search :: text != '' THEN (
			p.name ILIKE concat('%', @search, '%')
		)
		ELSE true
	END
  AND CASE
		WHEN @status IN (1,2)  THEN
			status = @status
		ELSE true
	END
  AND CASE
		WHEN @type :: integer > 0  THEN
			type = @type
		ELSE true
	END
ORDER BY 
  CASE 
        WHEN @order_by ::varchar = 'name asc' THEN p.name END ASC,
  CASE 
        WHEN @order_by = 'name desc' THEN p.name END DESC,
  CASE 
        WHEN @order_by = 'slug asc' THEN p.slug END ASC,
  CASE 
        WHEN @order_by = 'slug desc' THEN p.slug END DESC,
  CASE 
        WHEN @order_by = 'type asc' THEN type END ASC,
  CASE 
        WHEN @order_by = 'type desc' THEN type END DESC,
  CASE 
        WHEN @order_by = 'price asc' THEN price END ASC,
  CASE 
        WHEN @order_by = 'price desc' THEN price END DESC,
  CASE 
        WHEN @order_by = 'countInStock asc' THEN "countInStock" END ASC,
  CASE 
        WHEN @order_by = 'countInStock desc' THEN "countInStock" END DESC,
  CASE 
        WHEN @order_by = 'status asc' THEN status END ASC,
  CASE 
        WHEN @order_by = 'status desc' THEN status END DESC,
  CASE 
        WHEN @order_by = 'created_date asc' THEN p.create_at END ASC,
  CASE 
        WHEN @order_by = 'created_date desc' THEN p.create_at END DESC
LIMIT $1
OFFSET $2;

-- name: GetAllProductPublic :many
SELECT p.*,COUNT(l."user_id") AS "totalLikes",COUNT(p."_id") OVER() AS "totalCount",
CASE WHEN COUNT(l."user_id") > 0 THEN json_agg(l."user_id") ELSE '[]'::json END AS "likedBy",
CASE WHEN COUNT(v."user_id") > 0 THEN json_agg(v."user_id") ELSE '[]'::json END AS "uniqueViews"
FROM "Product" p
LEFT JOIN "Product_liked" l ON l."product_id" = p."_id"
LEFT JOIN "Product_UniqueView" v ON v."product_id" = p."_id"
WHERE 
  CASE
		WHEN @search :: text != '' THEN (
			p.name ILIKE concat('%', @search, '%')
		)
		ELSE true
	END
  AND CASE
		WHEN @status IN (1,2)  THEN
			status = @status
		ELSE true
	END
  AND CASE
		WHEN @type :: integer > 0  THEN
			type = @type
		ELSE true
	END
  AND CASE
		WHEN @minPrice :: integer > 0  THEN
			price >= @minPrice
		ELSE true
	END
  AND CASE
		WHEN @maxPrice :: integer > 0  THEN
			price <= @maxPrice
		ELSE true
	END
GROUP BY p."_id"
ORDER BY create_at ASC
LIMIT $1
OFFSET $2;


-- name: UpdateViewProduct :exec
UPDATE "Product" SET views = $1
WHERE "_id" = $2;

-- name: DeleteProductById :exec
DELETE FROM "Product"
WHERE "_id" = $1;

-- name: DeleteLikedProductByUserId :exec
DELETE FROM "Product_liked"
WHERE user_id = $1;

-- name: DeleteManyProductsByIds :exec
DELETE FROM "Product"
WHERE "_id" = ANY($1::bigint[]);