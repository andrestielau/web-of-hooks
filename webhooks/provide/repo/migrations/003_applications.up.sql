CREATE TYPE webhooks.new_application AS (
    id UUID,
    name TEXT,
    tenant_id TEXT,
    rate_limit INT,
    metadata JSONB
);

CREATE TABLE webhooks.application (
    id SERIAL PRIMARY KEY,
    name TEXT,
    uid UUID NOT NULL,
    tenant_id TEXT NOT NULL UNIQUE,
    rate_limit INT NOT NULL DEFAULT 0,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);