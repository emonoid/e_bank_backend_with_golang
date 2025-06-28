-- name: CreateEntry :one
INSERT INTO entries (
  -- id,
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
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateEntry :one
UPDATE entries
  set account_id = $2,
  amount = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :one
DELETE FROM entries
WHERE id = $1
RETURNING *;