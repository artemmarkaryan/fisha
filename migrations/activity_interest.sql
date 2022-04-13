create table if not exists "activity_interest"
(
    activity_id     int    not null,
    constraint activity_interest__activity_id_fk foreign key (activity_id) references "activity" (id) on delete cascade on update cascade,

    interest_id int    not null,
    constraint activity_interest__interest_id_fk foreign key (interest_id) references "interest" (id) on delete cascade on update cascade,

    rank        float4 not null default 0,

    created_at  timestamp       default current_timestamp,
    updated_at  timestamp       default current_timestamp
)