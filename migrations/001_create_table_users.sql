-- +goose Up
create table if not exists users  (
    id bigint primary key,
    chat_id bigint,
    username text,
    first_name text,
    last_name text,
    role text,
    token text,
    created_at timestamp,
    changed_at timestamp,
    deleted_at timestamp
);

-- +goose Down
drop table if exists users;