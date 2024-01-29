// Code generated by pggen. DO NOT EDIT.

package queries

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"time"
)

const createEventTypesSQL = `INSERT INTO webhooks.event_type (
    key
) 
SELECT 
    u.key
FROM unnest($1::webhooks.new_event_type[]) u
ON CONFLICT DO NOTHING
RETURNING 
    id,
    uid,
    key,
    created_at;`

type CreateEventTypesRow struct {
	ID        int32     `json:"id"`
	Uid       string    `json:"uid"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateEventTypes implements Querier.CreateEventTypes.
func (q *DBQuerier) CreateEventTypes(ctx context.Context, eventTypes []NewEventType) ([]CreateEventTypesRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "CreateEventTypes")
	rows, err := q.conn.Query(ctx, createEventTypesSQL, q.types.newNewEventTypeArrayInit(eventTypes))
	if err != nil {
		return nil, fmt.Errorf("query CreateEventTypes: %w", err)
	}
	defer rows.Close()
	items := []CreateEventTypesRow{}
	for rows.Next() {
		var item CreateEventTypesRow
		if err := rows.Scan(&item.ID, &item.Uid, &item.Key, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan CreateEventTypes row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close CreateEventTypes rows: %w", err)
	}
	return items, err
}

const deleteEventTypesSQL = `DELETE FROM webhooks.event_type WHERE key = ANY($1::TEXT[]);`

// DeleteEventTypes implements Querier.DeleteEventTypes.
func (q *DBQuerier) DeleteEventTypes(ctx context.Context, keys []string) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteEventTypes")
	cmdTag, err := q.conn.Exec(ctx, deleteEventTypesSQL, keys)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query DeleteEventTypes: %w", err)
	}
	return cmdTag, err
}

const listEventTypesSQL = `SELECT
    id,
    uid,
    key,
    created_at
FROM webhooks.event_type
ORDER BY uid
LIMIT $1
OFFSET $2;`

type ListEventTypesRow struct {
	ID        *int32    `json:"id"`
	Uid       string    `json:"uid"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
}

// ListEventTypes implements Querier.ListEventTypes.
func (q *DBQuerier) ListEventTypes(ctx context.Context, limit int, offset int) ([]ListEventTypesRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListEventTypes")
	rows, err := q.conn.Query(ctx, listEventTypesSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query ListEventTypes: %w", err)
	}
	defer rows.Close()
	items := []ListEventTypesRow{}
	for rows.Next() {
		var item ListEventTypesRow
		if err := rows.Scan(&item.ID, &item.Uid, &item.Key, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan ListEventTypes row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListEventTypes rows: %w", err)
	}
	return items, err
}
