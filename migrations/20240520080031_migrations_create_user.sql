-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         serial primary key,
    username   text      not null,
    password   text      not null,
    email      text      not null,
    created_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
