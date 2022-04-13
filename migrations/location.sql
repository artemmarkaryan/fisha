create table if not exists "location"
(
    id         serial,
    constraint location_id_pk primary key (id),
    name       text unique not null,
    lon        float       not null,
    lat        float       not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
