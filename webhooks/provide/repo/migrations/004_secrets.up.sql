CREATE TYPE webhooks.new_secret AS (
    tenant_id TEXT,
    id UUID
);
CREATE TABLE webhooks.secret (
    uid UUID UNIQUE NOT NULL,
    id SERIAL PRIMARY KEY,
    tenant_id TEXT,
    value TEXT
);