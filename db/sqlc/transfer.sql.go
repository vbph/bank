// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
insert into transfers (from_id, to_id, amount)
values ($1, $2, $3)
returning id, from_id, to_id, amount, created_at
`

type CreateTransferParams struct {
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromID, arg.ToID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromID,
		&i.ToID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const readAllTransfers = `-- name: ReadAllTransfers :many
select id, from_id, to_id, amount, created_at from transfers
where from_id = $1 or to_id = $2
limit $3
offset $4
`

type ReadAllTransfersParams struct {
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ReadAllTransfers(ctx context.Context, arg ReadAllTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, readAllTransfers,
		arg.FromID,
		arg.ToID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromID,
			&i.ToID,
			&i.Amount,
			&i.CreatedAt,
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

const readTransfer = `-- name: ReadTransfer :one
select id, from_id, to_id, amount, created_at from transfers
where id = $1
limit 1
`

func (q *Queries) ReadTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, readTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromID,
		&i.ToID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}