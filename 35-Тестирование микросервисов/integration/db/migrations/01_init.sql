-- +goose Up

create table if not exists items (
    id bigserial PRIMARY KEY,
    name text unique,
    description text,
    created_at timestamptz,
    updated_at timestamptz
);

-- +goose Down
