-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS player_wallet (
    tg INT PRIMARY KEY,
    credits INT NOT NULL,
    gold INT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS player_wallet;
-- +goose StatementEnd
