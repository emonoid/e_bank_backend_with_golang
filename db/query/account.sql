-- name: CreateAccount :one
INSERT INTO accounts (
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
ORDER BY id;

-- name: UpdateAccount :exec
UPDATE accounts
  set owner_name = $1,
  balance = $2,
  currency = $3
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;