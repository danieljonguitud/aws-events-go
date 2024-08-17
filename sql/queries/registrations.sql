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
