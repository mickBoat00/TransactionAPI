-- +goose Up
CREATE TABLE transactions (
  id UUID PRIMARY KEY NOT NULL,
  currency_id UUID NOT NULL,
  category_id UUID NOT NULL,
  amount DECIMAL(10,2) NOT NULL,
  date TIMESTAMP NOT NULL,
  user_id UUID NOT NULL,
  createdAt TIMESTAMP NOT NULL,
  updatedAt TIMESTAMP NOT NULL,
  CONSTRAINT fk_currency FOREIGN KEY (currency_id) REFERENCES currencies(id) ON DELETE CASCADE,
  CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE transactions;