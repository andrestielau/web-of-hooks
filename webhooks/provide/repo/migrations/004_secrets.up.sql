CREATE TYPE webhooks.new_secret AS (
    application_id INT,
    value TEXT
);
CREATE TABLE webhooks.secret (
    id SERIAL PRIMARY KEY,
    uid UUID UNIQUE NOT NULL DEFAULT generate_ulid(),
    application_id INT NOT NULL REFERENCES webhooks.application(id),
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);