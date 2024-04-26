-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id        bigserial PRIMARY KEY,
    username  varchar(25) NOT NULL,
    firstname varchar(25) NOT NULL,
    email     varchar(50) NOT NULL,
    password  varchar(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS;
-- +goose StatementEnd
