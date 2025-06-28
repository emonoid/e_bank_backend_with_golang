-- name: CreateAccount :one
INSERT INTO accounts (
  -- id,
  owner_name,
  balance,
  currency,
  created_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
  set owner_name = $2,
  balance = $3,
  currency = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :one
DELETE FROM accounts
WHERE id = $1
RETURNING *;