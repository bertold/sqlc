// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package db

import (
	"context"
)

const updateXWithY = `-- name: UpdateXWithY :exec
UPDATE x INNER JOIN y ON y.a = x.a SET x.b = y.b
`

func (q *Queries) UpdateXWithY(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateXWithY)
	return err
}
