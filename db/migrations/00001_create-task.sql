-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE task (
    id serial NOT NULL,
    name varchar(120) NOT NULL,
    updated_at timestamp with time zone DEFAULT now(),
    created_at timestamp with time zone DEFAULT now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE task;