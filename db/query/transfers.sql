-- name: CreateTransfer :one
INSERT INTO transfers (
  -- id,
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

-- name: GetTransferForUpdate :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers
  set from_account_id = $2,
  to_account_id = $3,
  amount = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :one
DELETE FROM transfers
WHERE id = $1
RETURNING *;