-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS exercises
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(128) NOT NULL,
    description TEXT         NOT NULL,
    duration    BIGINT,
    user_id     INT,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS exercises;
-- +goose StatementEnd
