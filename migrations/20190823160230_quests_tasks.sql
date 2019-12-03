-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS quests_tasks_seq;
CREATE TABLE IF NOT EXISTS quests_tasks (
    id INT PRIMARY KEY DEFAULT nextval('quests_tasks_seq'),
    quest INT NOT NULL,
    tasktype INT NOT NULL,
    taskvalue INT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS quests_tasks;
DROP SEQUENCE IF EXISTS quests_tasks_seq;
-- +goose StatementEnd
