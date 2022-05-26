-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE orders (
  order_id BIGSERIAL PRIMARY KEY,
<<<<<<< HEAD
  order_price FLOAT NOT NULL,
=======
  order_price text NOT NULL,
  idr text NOT NULL,
  btc text NOT NULL,
>>>>>>> test
  order_type text NOT NULL,
  order_crypto text NOT NULL,
  order_status text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);