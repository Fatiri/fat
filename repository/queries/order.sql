-- name: GetOrder :one
SELECT * FROM orders
WHERE order_id = $1 LIMIT 1;

-- name: ListOrder :many
SELECT * FROM orders
LIMIT $1
OFFSET $2;

-- name: CreateOrder :one
INSERT INTO orders (
  order_id, order_price, created_at, updated_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_id = $1;