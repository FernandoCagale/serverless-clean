-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO task (name, updated_at, created_at) VALUES ('Task 1', now(), now());
INSERT INTO task (name, updated_at, created_at) VALUES ('Task 2', now(), now());
INSERT INTO task (name, updated_at, created_at) VALUES ('Task 3', now(), now());

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DELETE FROM task;