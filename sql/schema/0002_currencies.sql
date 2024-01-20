-- +goose Up
CREATE TABLE currencies (
  id UUID PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  code VARCHAR(3) NOT NULL,
  createdAt TIMESTAMP NOT NULL,
  updatedAt TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE currencies;