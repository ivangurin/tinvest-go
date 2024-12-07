-- +goose Up
create table if not exists users  (
    id bigint primary key,
    username text,
    first_name text,
    last_name text,
    token text,
    created_at timestamp,
    changed_at timestamp,
    deleted_at timestamp,
    chat_id bigint,
    role text
);

-- +goose Down
drop table if exists users;