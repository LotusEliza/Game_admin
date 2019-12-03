-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS players (
    tg INT PRIMARY KEY,
    chat INT NOT NULL,
    name STRING NOT NULL,
    story INT NOT NULL,
    faction INT NOT NULL,
    referrer INT NOT NULL,
    location INT NOT NULL,
    posx INT NOT NULL,
    posy INT,
    hp INT NOT NULL,

    equipw INT,
    equipa INT,
    equipb INT,
    equipc INT,

    air INT,

    registered TIMESTAMP NOT NULL,
    lastactive TIMESTAMP NOT NULL,

    INDEX registered (registered),
    INDEX lastactive (lastactive)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS players;
-- +goose StatementEnd