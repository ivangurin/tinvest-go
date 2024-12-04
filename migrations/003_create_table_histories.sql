-- +goose Up
create table if not exists histories (
    id text primary key,
    created_at timestamp,
    user_id bigint,
    command text
);

-- +goose Down
drop table if exists histories;