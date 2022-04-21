create table if not exists "interest"
(
    id         serial,
    constraint interest__id_pk primary key (id),
    name       text not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
