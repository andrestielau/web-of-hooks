CREATE TYPE webhooks.new_event_type AS (
    key TEXT
);

CREATE TABLE webhooks.event_type (
    id SERIAL PRIMARY KEY,
    key TEXT NOT NULL UNIQUE,
    uid UUID NOT NULL UNIQUE DEFAULT generate_ulid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);