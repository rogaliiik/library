CREATE TABLE users
(
    id         serial primary key,
    username   text      not null,
    password   text      not null,
    email      text,
    created_at timestamp not null default now()
);