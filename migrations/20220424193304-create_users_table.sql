
-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id SERIAL NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT,
    password TEXT NOT NULL,
    date_created TIMESTAMPTZ NOT NULL,
    date_updated TIMESTAMPTZ
);


-- +migrate Down
DROP TABLE IF EXISTS users;
