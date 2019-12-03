-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS player_inventory (
    tg INT NOT NULL,
    itemtype STRING(20) NOT NULL,
    itemvalue INT NOT NULL,
    INDEX tg (tg)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS player_inventory;
-- +goose StatementEnd


-- CREATE SEQUENCE IF NOT EXISTS iventory_seq;
-- CREATE TABLE IF NOT EXISTS player_inventory (
--     id INT PRIMARY KEY DEFAULT nextval('iventory_seq'),
--     tg INT NOT NULL,
--     itemid INT NOT NULL,
--     count INT NOT NULL
-- );
