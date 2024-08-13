-- name: CreateEvent :one
INSERT INTO events (
    name, description, location, dateTime, user_id
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;
