
-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id SERIAL,
    username TEXT NOT NULL,
    email TEXT,
    password TEXT NOT NULL,
    "dateCreated" DATE NOT NULL,
    "dateUpdated" DATE,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);


-- +migrate Down
DROP TABLE IF EXISTS users;