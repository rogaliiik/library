-- +goose Up
-- +goose StatementBegin
CREATE TABLE books
(
    id         serial primary key,
    user_id    int,
    name       text      not null,
    content    text      not null,
    author     text      not null,
    created_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
