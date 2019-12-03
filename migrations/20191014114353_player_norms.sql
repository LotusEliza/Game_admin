-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- игрок int
-- дата Date
-- ресурс int
-- количество int

CREATE TABLE IF NOT EXISTS player_norms (
    tg INT,
    date DATE NOT NULL,
    resource INT NOT NULL,
    amount INT NOT NULL,
    PRIMARY KEY (tg, date, resource)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS player_norms;
-- +goose StatementEnd
