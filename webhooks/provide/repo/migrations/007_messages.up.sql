CREATE TYPE webhooks.new_message AS (
    application_id UUID,
    event_type_id UUID,
    event_id TEXT,
    payload TEXT
);
CREATE TABLE webhooks.message (
    id SERIAL PRIMARY KEY,
    application_id INT REFERENCES webhooks.application(id),
    event_type_id INT REFERENCES webhooks.event_type(id),
    uid UUID UNIQUE NOT NULL DEFAULT generate_ulid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    event_id TEXT UNIQUE,
    payload TEXT
);
CREATE TABLE webhooks.message_attempt (
    id SERIAL PRIMARY KEY,
    uid UUID UNIQUE NOT NULL DEFAULT generate_ulid(), --todo make mandatory for idempotency
    endpoint_id INT REFERENCES webhooks.endpoint(id),
    message_id INT REFERENCES webhooks.message(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    status INT NOT NULL DEFAULT 0,
    response_status INT,
    response TEXT
);
CREATE TYPE webhooks.message_details AS (
    message webhooks.message,
    attempts webhooks.message_attempt[]
);
