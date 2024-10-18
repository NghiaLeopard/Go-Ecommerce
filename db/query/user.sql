-- name: CreateUserAdmin :one
INSERT INTO "Users" (
  email,password,"firstName","lastName","middleName",avatar,address,"phoneNumber",role,city
) VALUES (
  $1,$2,$3,$4,$5,$6,$7,$8,$9,$10
)
RETURNING *;

-- name: GetUserAdminById :one
SELECT 
  u."_id",
  COALESCE(u."firstName", '') AS "firstName",
  COALESCE(u."lastName", '') AS "lastName",
  COALESCE(u."address", '') AS "address",
  COALESCE(u."middleName", '') AS "middleName",
  COALESCE(u."email", '') AS "email",
  COALESCE(u."avatar", '') AS "avatar",
  COALESCE(u."phoneNumber", '') AS "phoneNumber",
  COALESCE(u."status", u."status") AS "status",
  COALESCE(u."userType", u."userType") AS "userType",
  json_build_object('_id', r."_id",'name', r.name) AS "role",
  COALESCE(c."_id", 0) AS "city",
  COUNT(u."_id") OVER() AS "totalCount" 
FROM "Users" u
LEFT JOIN "Role" r ON r."_id" = u."role"
LEFT JOIN "City" c ON c."_id" = u."city"
WHERE u."_id" = $1 LIMIT 1;

-- name: GetUserAdminByEmail :one
SELECT * FROM "Users"
WHERE email = sqlc.arg(email) LIMIT 1;

-- name: ListUserAdmin :many
SELECT 
  u."_id",
  COALESCE(u."firstName", '') AS "firstName",
  COALESCE(u."lastName", '') AS "lastName",
  COALESCE(u."middleName", '') AS "middleName",
  COALESCE(u."email", '') AS "email",
  COALESCE(u."phoneNumber", '') AS "phoneNumber",
  COALESCE(u."status", u."status") AS "status",
  COALESCE(u."userType", u."userType") AS "userType",
  json_build_object('_id', r."_id",'name', r.name) AS "role",
  json_build_object('_id', c."_id",'name', c.name) AS "city",
  COUNT(u."_id") OVER() AS "totalCount" 
FROM "Users" u
LEFT JOIN "Role" r ON r."_id" = u."role"
LEFT JOIN "City" c ON c."_id" = u."city"
WHERE  @search ::text = '' or u."email" ILIKE concat('%',@search,'%')
ORDER BY 
  CASE 
        WHEN @order_by = 'createdAt asc' THEN u.create_at END ASC,
  CASE 
        WHEN @order_by = 'createdAt desc' THEN u.create_at END DESC
LIMIT NULLIF(@limit_opt :: int, 0)
OFFSET NULLIF(@offset_opt :: int, 0);


-- name: UpdateUserAdmin :one
UPDATE "Users" SET "firstName" = $1,"lastName" = $2,"middleName" = $3,avatar = $4,address = $5,"phoneNumber" = $6,role = $7,city = $8,status = $9,update_at = NOW()
WHERE "_id" = $10
RETURNING *;

-- name: DeleteUserAdminById :exec
DELETE FROM "Users"
WHERE "_id" = $1;

-- name: DeleteManyUserAdminByIds :exec
DELETE FROM "Users"
WHERE "_id" = ANY($1::bigint[]);
