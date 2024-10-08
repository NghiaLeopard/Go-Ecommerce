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
RETURNING id, name, permission, create_at, update_at
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
RETURNING id, name, permission, create_at, update_at
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
WHERE id = ANY($1::bigint[])
`

func (q *Queries) DeleteManyRolesByIds(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.ExecContext(ctx, deleteManyRolesByIds, pq.Array(dollar_1))
	return err
}

const deleteRoleById = `-- name: DeleteRoleById :exec
DELETE FROM "Role"
WHERE id = $1
`

func (q *Queries) DeleteRoleById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteRoleById, id)
	return err
}

const getRoleById = `-- name: GetRoleById :one
SELECT id, name, permission, create_at, update_at FROM "Role"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRoleById(ctx context.Context, id int64) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRoleById, id)
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

const getRoleByName = `-- name: GetRoleByName :one
SELECT id, name, permission, create_at, update_at FROM "Role"
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
SELECT id, name, permission, create_at, update_at,COUNT("Role".id) OVER() AS "totalCount" FROM "Role"
WHERE  $3 ::text = '' or name ILIKE concat('%',$3,'%')
LIMIT $1
OFFSET $2
`

type ListRoleParams struct {
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
	Search string `json:"search"`
}

type ListRoleRow struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Permission []string  `json:"permission"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
	TotalCount int64     `json:"totalCount"`
}

func (q *Queries) ListRole(ctx context.Context, arg ListRoleParams) ([]ListRoleRow, error) {
	rows, err := q.db.QueryContext(ctx, listRole, arg.Limit, arg.Offset, arg.Search)
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
WHERE id = $3
RETURNING id, name, permission, create_at, update_at
`

type UpdateRoleParams struct {
	Name       string   `json:"name"`
	Permission []string `json:"permission"`
	ID         int64    `json:"id"`
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
