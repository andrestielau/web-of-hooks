CREATE TABLE webhooks.statistics(
    endpoint_id INT REFERENCES webhooks.endpoint(id),
    successes INT NOT NULL DEFAULT 0,
    failures INT NOT NULL DEFAULT 0
);