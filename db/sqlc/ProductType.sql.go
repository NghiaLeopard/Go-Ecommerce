// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ProductType.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createProductType = `-- name: CreateProductType :one
INSERT INTO "Product_Type" (
  name,slug
) VALUES (
  $1, $2
)
RETURNING id, name, slug, create_at, update_at
`

type CreateProductTypeParams struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (q *Queries) CreateProductType(ctx context.Context, arg CreateProductTypeParams) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, createProductType, arg.Name, arg.Slug)
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteManyProductTypesByIds = `-- name: DeleteManyProductTypesByIds :exec
DELETE FROM "Product_Type"
WHERE id = ANY($1::bigint[])
`

func (q *Queries) DeleteManyProductTypesByIds(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.ExecContext(ctx, deleteManyProductTypesByIds, pq.Array(dollar_1))
	return err
}

const deleteProductTypeById = `-- name: DeleteProductTypeById :exec
DELETE FROM "Product_Type"
WHERE id = $1
`

func (q *Queries) DeleteProductTypeById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProductTypeById, id)
	return err
}

const getProductTypeById = `-- name: GetProductTypeById :one
SELECT id, name, slug, create_at, update_at FROM "Product_Type"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProductTypeById(ctx context.Context, id int64) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, getProductTypeById, id)
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const getProductTypeByName = `-- name: GetProductTypeByName :one
SELECT id, name, slug, create_at, update_at FROM "Product_Type"
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetProductTypeByName(ctx context.Context, name string) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, getProductTypeByName, name)
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const listProductType = `-- name: ListProductType :many
SELECT id, name, slug, create_at, update_at,COUNT("Product_Type".id) OVER() AS "totalCount" FROM "Product_Type"
WHERE  $3 ::text = '' or name ILIKE concat('%',$3,'%')
ORDER BY 
  CASE 
        WHEN $4 ::varchar = 'name asc' THEN name END ASC,
  CASE 
        WHEN $4 = 'name desc' THEN name END DESC,
  CASE 
        WHEN $4 = 'slug asc' THEN slug END ASC,
  CASE 
        WHEN $4 = 'slug desc' THEN slug END DESC,
  CASE 
        WHEN $4 = 'created_date asc' THEN create_at END ASC,
  CASE 
        WHEN $4 = 'created_date desc' THEN create_at END DESC
LIMIT $1
OFFSET $2
`

type ListProductTypeParams struct {
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
	Search  string `json:"search"`
	OrderBy string `json:"order_by"`
}

type ListProductTypeRow struct {
	ID         int64        `json:"id"`
	Name       string       `json:"name"`
	Slug       string       `json:"slug"`
	CreateAt   time.Time    `json:"create_at"`
	UpdateAt   sql.NullTime `json:"update_at"`
	TotalCount int64        `json:"totalCount"`
}

func (q *Queries) ListProductType(ctx context.Context, arg ListProductTypeParams) ([]ListProductTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, listProductType,
		arg.Limit,
		arg.Offset,
		arg.Search,
		arg.OrderBy,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductTypeRow{}
	for rows.Next() {
		var i ListProductTypeRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.CreateAt,
			&i.UpdateAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProductType = `-- name: UpdateProductType :one
UPDATE "Product_Type" SET name = $1,slug = $2,update_at = NOW()
WHERE id = $3
RETURNING id, name, slug, create_at, update_at
`

type UpdateProductTypeParams struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	ID   int64  `json:"id"`
}

func (q *Queries) UpdateProductType(ctx context.Context, arg UpdateProductTypeParams) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, updateProductType, arg.Name, arg.Slug, arg.ID)
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}
