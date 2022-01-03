-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE accounts (
  account_id BIGSERIAL PRIMARY KEY,
  username  text NOT NULL,
  password text NOT NULL,
  email text NOT NULL,
  type text NOT NULL
);