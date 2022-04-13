create table if not exists "user"
(
    id         serial,
    constraint user_id_pk primary key (id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

