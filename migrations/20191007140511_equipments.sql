-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE IF NOT EXISTS equip_seq;
CREATE TABLE IF NOT EXISTS equipments (
    id INT PRIMARY KEY DEFAULT nextval('equip_seq'),
    title STRING NOT NULL,
    type STRING NOT NULL,
    subtype STRING NOT NULL,
    buyprice INT NOT NULL,
    sellprice INT NOT NULL,
    reputation INT NOT NULL,

    damage INT NOT NULL,
    armor INT NOT NULL,
    air INT NOT NULL,
    mine INT NOT NULL,
    time INT NOT NULL,

    socketype INT NOT NULL,
    sockets INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS equipments;
DROP SEQUENCE IF EXISTS equip_seq;
-- +goose StatementEnd
