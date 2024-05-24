CREATE TABLE books
(
    id         serial primary key,
    user_id    int       not null,
    name       text      not null,
    content    text      not null,
    author     text      not null,
    created_at timestamp not null default now()
);