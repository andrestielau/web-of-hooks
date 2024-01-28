CREATE TYPE webhooks.new_endpoint AS (
    tenant_id TEXT,
    rate_limit INT,
    metadata JSONB,
    description TEXT,
    filterTypes TEXT[],
    channels TEXT[]
);

CREATE TABLE webhooks.endpoint (
    id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES webhooks.application(id),
    uid UUID UNIQUE NOT NULL DEFAULT generate_ulid(),
    rate_limit INT NOT NULL DEFAULT 0,
    metadata JSONB,
    disabled BOOLEAN DEFAULT false,
    description TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE TABLE webhooks.endpoint_secret (
    endpoint_id INT REFERENCES webhooks.endpoint(id),
    secret_id INT REFERENCES webhooks.secret(id),
    PRIMARY KEY(endpoint_id, secret_id)
);
CREATE TABLE webhooks.endpoint_filter (
    endpoint_id INT REFERENCES webhooks.endpoint(id),
    event_type_id INT REFERENCES webhooks.event_type(id),
    PRIMARY KEY(endpoint_id, event_type_id)
);
CREATE TABLE webhooks.endpoint_channel (
    endpoint_id INT REFERENCES webhooks.endpoint(id),
    channel_id INT REFERENCES webhooks.channel(id),
    PRIMARY KEY(endpoint_id, channel_id)
);