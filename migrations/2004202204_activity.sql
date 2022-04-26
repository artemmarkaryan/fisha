create table activity
(
    id         serial
        constraint activity__id_pk
            primary key,
    name       text not null
        unique,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    address    text,
    lon        numeric,
    lat        numeric
);

alter table activity
    owner to admin;
