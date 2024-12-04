-- +goose Up
create table if not exists currencies (
    id text,
    date timestamp,
    rate float,
    changed_at timestamp,
    PRIMARY KEY (id, date)
);

-- +goose Down
drop table if exists currencies;