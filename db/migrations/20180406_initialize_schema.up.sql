create table commands (
    id integer not null primary key,
    cmd text not null,
    title text not null,
    description text,
    url text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);