// Code generated by pggen. DO NOT EDIT.

package queries

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"time"
)

const createApplicationsSQL = `INSERT INTO webhooks.application (
    tenant_id,
    rate_limit,
    metadata
) 
SELECT 
    u.tenant_id,
    u.rate_limit,
    u.metadata
FROM unnest($1::webhooks.new_application[]) u
ON CONFLICT DO NOTHING
RETURNING 
    id,
    uid,
    tenant_id,
    rate_limit,
    metadata,
    created_at;`

type CreateApplicationsRow struct {
	ID        int32        `json:"id"`
	Uid       string       `json:"uid"`
	TenantID  string       `json:"tenant_id"`
	RateLimit int32        `json:"rate_limit"`
	Metadata  pgtype.JSONB `json:"metadata"`
	CreatedAt time.Time    `json:"created_at"`
}

// CreateApplications implements Querier.CreateApplications.
func (q *DBQuerier) CreateApplications(ctx context.Context, applications []NewApplication) ([]CreateApplicationsRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "CreateApplications")
	rows, err := q.conn.Query(ctx, createApplicationsSQL, q.types.newNewApplicationArrayInit(applications))
	if err != nil {
		return nil, fmt.Errorf("query CreateApplications: %w", err)
	}
	defer rows.Close()
	items := []CreateApplicationsRow{}
	for rows.Next() {
		var item CreateApplicationsRow
		if err := rows.Scan(&item.ID, &item.Uid, &item.TenantID, &item.RateLimit, &item.Metadata, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan CreateApplications row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close CreateApplications rows: %w", err)
	}
	return items, err
}

const listApplicationsSQL = `SELECT
    id,
    uid,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
FROM webhooks.application
ORDER BY uid
LIMIT $1
OFFSET $2;`

type ListApplicationsRow struct {
	ID        *int32       `json:"id"`
	Uid       string       `json:"uid"`
	TenantID  string       `json:"tenant_id"`
	RateLimit *int32       `json:"rate_limit"`
	Metadata  pgtype.JSONB `json:"metadata"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// ListApplications implements Querier.ListApplications.
func (q *DBQuerier) ListApplications(ctx context.Context, limit int, offset int) ([]ListApplicationsRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListApplications")
	rows, err := q.conn.Query(ctx, listApplicationsSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query ListApplications: %w", err)
	}
	defer rows.Close()
	items := []ListApplicationsRow{}
	for rows.Next() {
		var item ListApplicationsRow
		if err := rows.Scan(&item.ID, &item.Uid, &item.TenantID, &item.RateLimit, &item.Metadata, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan ListApplications row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListApplications rows: %w", err)
	}
	return items, err
}