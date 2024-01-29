// Code generated by pggen. DO NOT EDIT.

package queries

import (
	"context"
	"fmt"
	"time"
)

const listMessagesSQL = `SELECT
    id,
    uid,
    created_at
FROM webhooks.message
ORDER BY uid
LIMIT $1
OFFSET $2;`

type ListMessagesRow struct {
	ID        *int32    `json:"id"`
	Uid       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
}

// ListMessages implements Querier.ListMessages.
func (q *DBQuerier) ListMessages(ctx context.Context, limit int, offset int) ([]ListMessagesRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListMessages")
	rows, err := q.conn.Query(ctx, listMessagesSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query ListMessages: %w", err)
	}
	defer rows.Close()
	items := []ListMessagesRow{}
	for rows.Next() {
		var item ListMessagesRow
		if err := rows.Scan(&item.ID, &item.Uid, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan ListMessages row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListMessages rows: %w", err)
	}
	return items, err
}
