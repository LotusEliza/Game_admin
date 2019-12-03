-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS player_quests (
    tg INT NOT NULL,
    quest INT NOT NULL,
    created TIMESTAMP NOT NULL,
    tasktype INT NOT NULL,
    taskvalue INT NOT NULL,
    INDEX tg (tg)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS player_quests;
-- +goose StatementEnd
