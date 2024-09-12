-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS exercises
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(128) NOT NULL,
    description TEXT         NOT NULL,
    duration    BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS exercises;
-- +goose StatementEnd
