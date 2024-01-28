CREATE TYPE webhooks.new_application AS (
    tenant_id TEXT,
    rate_limit INT,
    metadata JSONB
);

CREATE TABLE webhooks.application (
    id SERIAL PRIMARY KEY,
    tenant_id TEXT NOT NULL UNIQUE,
    uid UUID NOT NULL UNIQUE DEFAULT generate_ulid(),
    rate_limit INT NOT NULL DEFAULT 0,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);