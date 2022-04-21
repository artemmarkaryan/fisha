create table if not exists "activity"
(
    id          serial
        constraint activity__id_pk primary key,

    location_id int  not null,
    constraint activity__location_id_fk foreign key (location_id) references "location" (id) on delete cascade on update cascade,

    name        text not null unique,

    created_at  timestamp default current_timestamp,
    updated_at  timestamp default current_timestamp
);