-- +goose Up
CREATE TABLE sources (
     id SERIAL PRIMARY KEY,
     url VARCHAR(255) NOT NULL UNIQUE,
     kind VARCHAR(255) NOT NULL,
     last_update_at TIMESTAMP WITH TIME ZONE,
     created_at TIMESTAMP WITH TIME ZONE NOT NULL ,
     updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS sources;