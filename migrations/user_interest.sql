create table if not exists "user_interest"
(
    user_id     int    not null,
    constraint user_interest__user_id_fk foreign key (user_id) references "user" (id) on delete cascade on update cascade,

    interest_id int    not null,
    constraint user_interest__interest_id_fk foreign key (interest_id) references "interest" (id) on delete cascade on update cascade,

    rank        float4 not null default 0,

    created_at  timestamp       default current_timestamp,
    updated_at  timestamp       default current_timestamp
)