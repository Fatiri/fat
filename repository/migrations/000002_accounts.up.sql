-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE accounts (
  account_id UUID PRIMARY KEY,
  username  text NOT NULL,
  password text NOT NULL,
  email text NOT NULL,
  account_type text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);