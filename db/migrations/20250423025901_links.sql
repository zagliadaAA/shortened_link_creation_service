-- +goose Up
-- +goose StatementBegin
create table links(
    ID bigserial primary key,
    url varchar(200) not null unique,
    short_url varchar(40) not null unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table links;
-- +goose StatementEnd
