-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;


-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ListAccounts :many
SELECT * FROM accounts
order by id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
  set balance = $2
WHERE id = $1 RETURNING *;


-- name: AddToAccountBalance :one
UPDATE accounts
  set balance = sqlc.arg(amount) + balance
WHERE id = sqlc.arg(id) RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;