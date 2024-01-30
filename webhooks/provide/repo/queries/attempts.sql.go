// Code generated by pggen. DO NOT EDIT.

package queries

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"time"
)

const deleteAttemptsSQL = `DELETE FROM webhooks.message_attempt WHERE uid = ANY($1::UUID[]);`

// DeleteAttempts implements Querier.DeleteAttempts.
func (q *DBQuerier) DeleteAttempts(ctx context.Context, ids []string) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteAttempts")
	cmdTag, err := q.conn.Exec(ctx, deleteAttemptsSQL, ids)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query DeleteAttempts: %w", err)
	}
	return cmdTag, err
}

const getAttemptsSQL = `SELECT 
    id,
    uid,
    created_at
FROM webhooks.message_attempt
WHERE uid = ANY($1::uuid[]);`

type GetAttemptsRow struct {
	ID        *int32    `json:"id"`
	Uid       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
}

// GetAttempts implements Querier.GetAttempts.
func (q *DBQuerier) GetAttempts(ctx context.Context, ids []string) ([]GetAttemptsRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetAttempts")
	rows, err := q.conn.Query(ctx, getAttemptsSQL, ids)
	if err != nil {
		return nil, fmt.Errorf("query GetAttempts: %w", err)
	}
	defer rows.Close()
	items := []GetAttemptsRow{}
	for rows.Next() {
		var item GetAttemptsRow
		if err := rows.Scan(&item.ID, &item.Uid, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan GetAttempts row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close GetAttempts rows: %w", err)
	}
	return items, err
}

const listAttemptsSQL = `SELECT
    id,
    uid,
    created_at
FROM webhooks.message_attempt
ORDER BY uid
LIMIT $1
OFFSET $2;`

type ListAttemptsRow struct {
	ID        *int32    `json:"id"`
	Uid       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
}

// ListAttempts implements Querier.ListAttempts.
func (q *DBQuerier) ListAttempts(ctx context.Context, limit int, offset int) ([]ListAttemptsRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListAttempts")
	rows, err := q.conn.Query(ctx, listAttemptsSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query ListAttempts: %w", err)
	}
	defer rows.Close()
	items := []ListAttemptsRow{}
	for rows.Next() {
		var item ListAttemptsRow
		if err := rows.Scan(&item.ID, &item.Uid, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan ListAttempts row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListAttempts rows: %w", err)
	}
	return items, err
}
