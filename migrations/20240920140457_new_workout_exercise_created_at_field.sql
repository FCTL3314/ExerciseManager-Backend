-- +goose Up
-- +goose StatementBegin
ALTER TABLE workout_exercises ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE workout_exercises DROP COLUMN created_at;
-- +goose StatementEnd
