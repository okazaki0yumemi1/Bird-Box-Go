-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS birds
(
    id                  UUID        PRIMARY KEY,
    common_name         TEXT        NOT NULL,
    confidence          TEXT        NOT NULL,
    recoding_date       DATE        NULL,
    microphone_name     TEXT        NULL,
    created_at          TIMESTAMP   NOT NULL,
    updated_at          TIMESTAMP   NOT NULL,
    deleted_at          TIMESTAMP   NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS birds;