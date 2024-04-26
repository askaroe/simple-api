-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN lastname varchar(25) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN lastname;
-- +goose StatementEnd
