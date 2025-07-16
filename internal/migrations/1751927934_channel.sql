-- +goose Up
CREATE TABLE channel (
     id SERIAL PRIMARY KEY,
     title VARCHAR(255) NOT NULL,
     description VARCHAR(255) NOT NULL,
     link VARCHAR(255) NOT NULL
);

Alter TABLE records ADD channel_id integer REFERENCES channel (id);

-- +goose Down
ALTER TABLE records DROP COLUMN channel_id;

DROP TABLE IF EXISTS channel;