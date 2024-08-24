-- name: CreateUser :one
INSERT INTO users(
    email, password
) VALUES (
    ?, ?
)
RETURNING id, email;

-- name: GetUser :one
SELECT * 
FROM users
WHERE email = ?;
