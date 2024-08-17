-- name: CreateEvent :one
INSERT INTO events (
    name, description, location, dateTime, user_id
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: ListEvents :many
SELECT * FROM events;

-- name: GetEvent :one
SELECT * FROM events
WHERE id = ?;

-- name: UpdateEvent :one
UPDATE events
SET name = ?, description = ?, location = ?, dateTime = ?
WHERE id = ?
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = ?;

-- name: RegisterEventUser :one
INSERT INTO registrations (
    event_id, user_id
) VALUES (
    ?, ?
)
RETURNING *;

-- name: DeleteRegistration :exec
DELETE FROM registrations 
WHERE event_id = ? AND user_id = ?;
