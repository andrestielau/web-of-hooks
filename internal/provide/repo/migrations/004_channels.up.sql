
CREATE TYPE webhooks.new_channel AS (
    key TEXT
);


CREATE TABLE webhooks.channel(
    uid UUID NOT NULL UNIQUE DEFAULT generate_ulid(),
    id SERIAL PRIMARY KEY,
    key TEXT UNIQUE
);