// Code generated by pggen. DO NOT EDIT.

package queries

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"time"
)

const createApplicationsSQL = `INSERT INTO webhooks.application (
    name,
    tenant_id,
    rate_limit,
    metadata
) 
SELECT 
    u.name,
    u.tenant_id,
    u.rate_limit,
    u.metadata
FROM unnest($1::webhooks.new_application[]) u
ON CONFLICT DO NOTHING
RETURNING (
    id,
    uid,
    name,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
)::webhooks.application;`

// CreateApplications implements Querier.CreateApplications.
func (q *DBQuerier) CreateApplications(ctx context.Context, applications []NewApplication) ([]Application, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "CreateApplications")
	rows, err := q.conn.Query(ctx, createApplicationsSQL, q.types.newNewApplicationArrayInit(applications))
	if err != nil {
		return nil, fmt.Errorf("query CreateApplications: %w", err)
	}
	defer rows.Close()
	items := []Application{}
	rowRow := q.types.newApplication()
	for rows.Next() {
		var item Application
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan CreateApplications row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign CreateApplications row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close CreateApplications rows: %w", err)
	}
	return items, err
}

const deleteApplicationsSQL = `DELETE FROM webhooks.application WHERE uid = ANY($1::UUID[]);`

// DeleteApplications implements Querier.DeleteApplications.
func (q *DBQuerier) DeleteApplications(ctx context.Context, ids []string) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteApplications")
	cmdTag, err := q.conn.Exec(ctx, deleteApplicationsSQL, ids)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query DeleteApplications: %w", err)
	}
	return cmdTag, err
}

const getApplicationsSQL = `SELECT (
    id,
    uid,
    name,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
)::webhooks.application
FROM webhooks.application
WHERE uid = ANY($1::uuid[]);`

// GetApplications implements Querier.GetApplications.
func (q *DBQuerier) GetApplications(ctx context.Context, ids []string) ([]Application, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetApplications")
	rows, err := q.conn.Query(ctx, getApplicationsSQL, ids)
	if err != nil {
		return nil, fmt.Errorf("query GetApplications: %w", err)
	}
	defer rows.Close()
	items := []Application{}
	rowRow := q.types.newApplication()
	for rows.Next() {
		var item Application
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan GetApplications row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign GetApplications row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close GetApplications rows: %w", err)
	}
	return items, err
}

const listApplicationsSQL = `SELECT (
    id,
    uid,
    name,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
)::webhooks.application
FROM webhooks.application
WHERE created_at > $1 
ORDER BY created_at
LIMIT $2
OFFSET $3;`

type ListApplicationsParams struct {
	CreatedAfter time.Time `json:"created_after"`
	Limit        int       `json:"limit"`
	Offset       int       `json:"offset"`
}

// ListApplications implements Querier.ListApplications.
func (q *DBQuerier) ListApplications(ctx context.Context, params ListApplicationsParams) ([]Application, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListApplications")
	rows, err := q.conn.Query(ctx, listApplicationsSQL, params.CreatedAfter, params.Limit, params.Offset)
	if err != nil {
		return nil, fmt.Errorf("query ListApplications: %w", err)
	}
	defer rows.Close()
	items := []Application{}
	rowRow := q.types.newApplication()
	for rows.Next() {
		var item Application
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan ListApplications row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign ListApplications row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListApplications rows: %w", err)
	}
	return items, err
}
