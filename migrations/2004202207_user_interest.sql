create table user_interest
(
    user_id     integer             not null
        constraint user_interest__user_id_fk
            references "user"
            on update cascade on delete cascade,
    interest_id integer             not null
        constraint user_interest__interest_id_fk
            references interest
            on update cascade on delete cascade,
    rank        real      default 0 not null
        constraint user_interest__rank_boundaries
            check ((rank >= ('-1'::integer)::double precision) AND (rank <= (1)::double precision)),
    created_at  timestamp default CURRENT_TIMESTAMP,
    updated_at  timestamp default CURRENT_TIMESTAMP
);

alter table user_interest
    owner to admin;

