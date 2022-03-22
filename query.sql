
-- name: GetPerson :one
SELECT * FROM persons
WHERE id = $1 LIMIT 1;

-- name: ListPersons :many
SELECT * FROM persons
ORDER BY Email;

-- name: CreatePerson :one
INSERT INTO persons (
  FirstName, LastName, Email
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeletePerson :exec
DELETE FROM persons
WHERE id = $1;