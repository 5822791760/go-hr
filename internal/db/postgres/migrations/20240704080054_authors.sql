-- +goose Up
-- +goose StatementBegin
CREATE TABLE authors (
  id   SERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE authors;
-- +goose StatementEnd
