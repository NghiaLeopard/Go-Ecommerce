// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: Role.sql

package db

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const createRole = `-- name: CreateRole :one
INSERT INTO "Role" (
  name
) VALUES (
  $1
)
RETURNING _id, name, permission, create_at, update_at
`

func (q *Queries) CreateRole(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole, name)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		pq.Array(&i.Permission),
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const createRoleByDefault = `-- name: CreateRoleByDefault :one
INSERT INTO "Role" (
  name, permission
) VALUES (
  $1, $2
)
RETURNING _id, name, permission, create_at, update_at
`

type CreateRoleByDefaultParams struct {
	Name       string   `json:"name"`
	Permission []string `json:"permission"`
}

func (q *Queries) CreateRoleByDefault(ctx context.Context, arg CreateRoleByDefaultParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRoleByDefault, arg.Name, pq.Array(arg.Permission))
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		pq.Array(&i.Permission),
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteManyRolesByIds = `-- name: DeleteManyRolesByIds :exec
DELETE FROM "Role"
WHERE "_id" = ANY($1::bigint[])
`

func (q *Queries) DeleteManyRolesByIds(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.ExecContext(ctx, deleteManyRolesByIds, pq.Array(dollar_1))
	return err
}

const deleteRoleById = `-- name: DeleteRoleById :exec
DELETE FROM "Role"
WHERE "_id" = $1
`

func (q *Queries) DeleteRoleById(ctx context.Context, ID int64) error {
	_, err := q.db.ExecContext(ctx, deleteRoleById, ID)
	return err
}

const getRoleById = `-- name: GetRoleById :one
SELECT _id,name,coalesce(permission, ARRAY[]::TEXT[]) FROM "Role"
WHERE _id = $1 LIMIT 1
`

type GetRoleByIdRow struct {
	ID         int64    `json:"_id"`
	Name       string   `json:"name"`
	Permission []string `json:"permission"`
}

func (q *Queries) GetRoleById(ctx context.Context, ID int64) (GetRoleByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getRoleById, ID)
	var i GetRoleByIdRow
	err := row.Scan(&i.ID, &i.Name, pq.Array(&i.Permission))
	return i, err
}

const getRoleByName = `-- name: GetRoleByName :one
SELECT _id, name, permission, create_at, update_at FROM "Role"
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetRoleByName(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRoleByName, name)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		pq.Array(&i.Permission),
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const listRole = `-- name: ListRole :many
SELECT _id, name, permission, create_at, update_at,COUNT("Role"."_id") OVER() AS "totalCount" FROM "Role"
WHERE  $1 ::text = '' or name ILIKE concat('%',$1,'%')
LIMIT NULLIF($3 :: int, 0)
OFFSET NULLIF($2 :: int, 0)
`

type ListRoleParams struct {
	Search    string `json:"search"`
	OffsetOpt int32  `json:"offset_opt"`
	LimitOpt  int32  `json:"limit_opt"`
}

type ListRoleRow struct {
	ID         int64     `json:"_id"`
	Name       string    `json:"name"`
	Permission []string  `json:"permission"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
	TotalCount int64     `json:"totalCount"`
}

func (q *Queries) ListRole(ctx context.Context, arg ListRoleParams) ([]ListRoleRow, error) {
	rows, err := q.db.QueryContext(ctx, listRole, arg.Search, arg.OffsetOpt, arg.LimitOpt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListRoleRow{}
	for rows.Next() {
		var i ListRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			pq.Array(&i.Permission),
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

const updateRole = `-- name: UpdateRole :one
UPDATE "Role" SET name = $1,permission = $2,update_at = NOW()
WHERE "_id" = $3
RETURNING _id, name, permission, create_at, update_at
`

type UpdateRoleParams struct {
	Name       string   `json:"name"`
	Permission []string `json:"permission"`
	ID         int64    `json:"_id"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, updateRole, arg.Name, pq.Array(arg.Permission), arg.ID)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		pq.Array(&i.Permission),
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}
