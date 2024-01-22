-- +goose Up
CREATE TABLE categories (
  id UUID PRIMARY KEY NOT NULL,
  name VARCHAR(20) NOT NULL,
  user_id UUID NOT NULL,
  createdAt TIMESTAMP NOT NULL,
  updatedAt TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE categories;