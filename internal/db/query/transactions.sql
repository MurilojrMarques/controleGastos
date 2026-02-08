-- name: CreateTransaction :one
INSERT INTO transactions (
  title, amount, type, due_date
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY due_date DESC;

-- name: UpdateTransaction :one
UPDATE transactions
SET 
  title = $2,
  amount = $3,
  type = $4,
  due_date = $5
WHERE id = $1
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;