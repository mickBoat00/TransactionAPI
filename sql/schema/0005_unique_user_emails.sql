-- +goose Up
ALTER TABLE users
ADD UNIQUE (email);

-- +goose Down
ALTER TABLE users DROP CONSTRAINT users_email_key;