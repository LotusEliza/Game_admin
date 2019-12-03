-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS player_duty (
    tg INT,
    timestart TIMESTAMP NOT NULL,
    timeend TIMESTAMP NOT NULL,
    duty INT NOT NULL,
    INDEX tg (tg),
    INDEX timeend (timeend)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS player_duty;
-- +goose StatementEnd
