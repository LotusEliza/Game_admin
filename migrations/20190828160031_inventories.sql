-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS items_seq;
CREATE TABLE IF NOT EXISTS items (
    id INT PRIMARY KEY DEFAULT nextval('quests_seq'),
    title STRING NOT NULL,
    itemtype STRING(20) NOT NULL,
    p1 INT NOT NULL,
    p2 INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS items;
DROP SEQUENCE IF EXISTS items_seq;
-- +goose StatementEnd
