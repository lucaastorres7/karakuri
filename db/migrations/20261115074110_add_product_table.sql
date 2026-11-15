-- +goose Up
CREATE TABLE IF NOT EXISTS products (
  "id"          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "name"        VARCHAR(255) NOT NULL,
  "description" TEXT,
  "image"       VARCHAR(255),
  "price"       DECIMAL(10, 2) NOT NULL,
  "quantity"    INT NOT NULL,
  "created_at"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS products;