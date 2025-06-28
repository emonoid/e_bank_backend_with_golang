-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount,
  created_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id;

-- name: UpdateTransfer :exec
UPDATE transfers
  set from_account_id = $1,
  to_account_id = $2,
  amount = $3
WHERE id = $1;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;