-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS quests_rewards_seq;
CREATE TABLE IF NOT EXISTS quests_rewards (
    id INT PRIMARY KEY DEFAULT nextval('quests_rewards_seq'),
    quest INT NOT NULL,
    rewardtype INT NOT NULL,
    rewardvalue INT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS quests_rewards;
DROP SEQUENCE IF EXISTS quests_rewards_seq;
-- +goose StatementEnd
