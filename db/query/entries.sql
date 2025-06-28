-- name: CreateEntry :one
INSERT INTO entries (
  account_id,
  amount,
  created_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntry :many
SELECT * FROM entries
ORDER BY id;

-- name: UpdateEntry :exec
UPDATE entries
  set account_id = $1,
  amount = $2
WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;