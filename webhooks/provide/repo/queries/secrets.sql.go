// Code generated by pggen. DO NOT EDIT.

package queries

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"time"
)

const createSecretsSQL = `INSERT INTO webhooks.secret (
    value,
    application_id
) 
SELECT 
    u.value,
    a.id
FROM unnest($1::webhooks.new_secret[]) u
 JOIN webhooks.application a ON u.application_id = a.uid
ON CONFLICT DO NOTHING
RETURNING (
    id,
    uid,
    application_id,
    value,
    created_at,
    updated_at
)::webhooks.secret;`

// CreateSecrets implements Querier.CreateSecrets.
func (q *DBQuerier) CreateSecrets(ctx context.Context, secrets []NewSecret) ([]Secret, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "CreateSecrets")
	rows, err := q.conn.Query(ctx, createSecretsSQL, q.types.newNewSecretArrayInit(secrets))
	if err != nil {
		return nil, fmt.Errorf("query CreateSecrets: %w", err)
	}
	defer rows.Close()
	items := []Secret{}
	rowRow := q.types.newSecret()
	for rows.Next() {
		var item Secret
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan CreateSecrets row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign CreateSecrets row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close CreateSecrets rows: %w", err)
	}
	return items, err
}

const getSecretsSQL = `SELECT (
    id,
    uid,
    application_id,
    value,
    created_at,
    updated_at
)::webhooks.secret
FROM webhooks.secret
WHERE uid = ANY($1::uuid[]);`

// GetSecrets implements Querier.GetSecrets.
func (q *DBQuerier) GetSecrets(ctx context.Context, ids []string) ([]Secret, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetSecrets")
	rows, err := q.conn.Query(ctx, getSecretsSQL, ids)
	if err != nil {
		return nil, fmt.Errorf("query GetSecrets: %w", err)
	}
	defer rows.Close()
	items := []Secret{}
	rowRow := q.types.newSecret()
	for rows.Next() {
		var item Secret
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan GetSecrets row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign GetSecrets row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close GetSecrets rows: %w", err)
	}
	return items, err
}

const deleteSecretsSQL = `DELETE FROM webhooks.secret WHERE uid = ANY($1::UUID[]);`

// DeleteSecrets implements Querier.DeleteSecrets.
func (q *DBQuerier) DeleteSecrets(ctx context.Context, ids []string) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteSecrets")
	cmdTag, err := q.conn.Exec(ctx, deleteSecretsSQL, ids)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query DeleteSecrets: %w", err)
	}
	return cmdTag, err
}

const listSecretsSQL = `SELECT (
    id,
    uid,
    application_id,
    value,
    created_at,
    updated_at
)::webhooks.secret
FROM webhooks.secret
WHERE created_at > $1 
ORDER BY uid
LIMIT $2
OFFSET $3;`

type ListSecretsParams struct {
	CreatedAfter time.Time `json:"created_after"`
	Limit        int       `json:"limit"`
	Offset       int       `json:"offset"`
}

// ListSecrets implements Querier.ListSecrets.
func (q *DBQuerier) ListSecrets(ctx context.Context, params ListSecretsParams) ([]Secret, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListSecrets")
	rows, err := q.conn.Query(ctx, listSecretsSQL, params.CreatedAfter, params.Limit, params.Offset)
	if err != nil {
		return nil, fmt.Errorf("query ListSecrets: %w", err)
	}
	defer rows.Close()
	items := []Secret{}
	rowRow := q.types.newSecret()
	for rows.Next() {
		var item Secret
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan ListSecrets row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign ListSecrets row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListSecrets rows: %w", err)
	}
	return items, err
}

const updateSecretSQL = `UPDATE webhooks.secret SET value = $1 WHERE uid = $2;`

// UpdateSecret implements Querier.UpdateSecret.
func (q *DBQuerier) UpdateSecret(ctx context.Context, value string, uid string) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateSecret")
	cmdTag, err := q.conn.Exec(ctx, updateSecretSQL, value, uid)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query UpdateSecret: %w", err)
	}
	return cmdTag, err
}

const listApplicationSecretsSQL = `SELECT (
    s.id,
    s.uid,
    s.application_id,
    s.value,
    s.created_at,
    s.updated_at
)::webhooks.secret
FROM webhooks.secret s
JOIN webhooks.application ON s.application_id = webhooks.application.id
WHERE webhooks.application.uid = $1::uuid
ORDER BY s.uid;`

// ListApplicationSecrets implements Querier.ListApplicationSecrets.
func (q *DBQuerier) ListApplicationSecrets(ctx context.Context, applicationUid string) ([]Secret, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListApplicationSecrets")
	rows, err := q.conn.Query(ctx, listApplicationSecretsSQL, applicationUid)
	if err != nil {
		return nil, fmt.Errorf("query ListApplicationSecrets: %w", err)
	}
	defer rows.Close()
	items := []Secret{}
	rowRow := q.types.newSecret()
	for rows.Next() {
		var item Secret
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan ListApplicationSecrets row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign ListApplicationSecrets row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListApplicationSecrets rows: %w", err)
	}
	return items, err
}
