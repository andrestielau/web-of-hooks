-- SetLastSeen
-- name: SetLastSeen :exec
INSERT INTO webhooks.worker (
    id
) VALUES (
    pggen.arg('id')
) ON CONFLICT(id)
    DO UPDATE SET 
        last_seen_at = now();

-- DequeueAttempts dequeues pending attempts for app with a hacked-up consistent hashing
-- name: DequeueAttempts :many
WITH last_seen AS (
    INSERT INTO webhooks.worker (
        id
    ) VALUES (
        pggen.arg('id')
    ) ON CONFLICT(id)
        DO UPDATE SET 
            last_seen_at = now()
), parts AS (
    SELECT 
        COUNT(*) as n, -- number of workers
        COUNT(first_seen_at < pggen.arg('start')) as i -- number of workers that started before caller
    FROM webhooks.worker
    WHERE last_seen_at = now() - interval '30 second' -- exclude idle workers
), app_ids AS ( -- applications for worker
    SELECT 
        id 
    FROM webhooks.application, parts
    WHERE MOD(id, parts.n) = parts.i
), selected AS (
    SELECT 
        e.application_id,
        a.endpoint_id,
        a.message_id,
        a.id as attempt_id
    FROM webhooks.message_attempt a
    INNER JOIN webhooks.endpoint e
        ON a.endpoint_id = e.id
    WHERE (status = 0 OR 
        (status = 1 AND a.updated_at < now() - interval '30 second')) -- repick old messages 
        AND e.application_id IN (SELECT id FROM app_ids) -- damn sql, why cant it be just app_ids?
    LIMIT pggen.arg('limit') 
), dequeue AS (
    UPDATE webhooks.message_attempt a
    SET status = 1, updated_at = now()
    FROM selected s
    WHERE a.id = s.attempt_id
)
SELECT 
    application_id,
    endpoint_id,
    message_id,
    attempt_id
FROM selected;
