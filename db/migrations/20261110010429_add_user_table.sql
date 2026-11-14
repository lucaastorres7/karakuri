-- +goose Up
CREATE TYPE gender AS ENUM (
  'male',
  'female',
  'other'
);

CREATE TABLE IF NOT EXISTS users (
  "id"         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "name"       VARCHAR(255) NOT NULL,
  "email"      VARCHAR(255) NOT NULL UNIQUE,
  "password"   VARCHAR(255) NOT NULL,
  "gender"     gender,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS gender;