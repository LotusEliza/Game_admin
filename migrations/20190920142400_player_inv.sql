-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS iventory_seq;
CREATE TABLE IF NOT EXISTS player_inv (
    id INT PRIMARY KEY DEFAULT nextval('iventory_seq'),
    tg INT NOT NULL,
    itemid INT NOT NULL,
    count INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS player_inv;
-- +goose StatementEnd