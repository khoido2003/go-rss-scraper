-- goose postgres postgres://username:password@localhost:5432/rss-scraper-db up


-- +goose Up

CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;