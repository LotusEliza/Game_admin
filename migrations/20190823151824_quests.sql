-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS quests_seq;
CREATE TABLE IF NOT EXISTS quests (
    id INT PRIMARY KEY DEFAULT nextval('quests_seq'),
    title STRING NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS quests;
DROP SEQUENCE IF EXISTS quests_seq;
-- +goose StatementEnd
