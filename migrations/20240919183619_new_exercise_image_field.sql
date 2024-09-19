-- +goose Up
-- +goose StatementBegin
ALTER TABLE exercises ADD COLUMN image TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE exercises DROP COLUMN image;
-- +goose StatementEnd
