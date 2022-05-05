create table recommendations
(
    user_id     integer
        constraint recommendations_user_id_fk
            references "user"
            on update cascade on delete cascade,
    activity_id integer
        constraint recommendations_activity_id_fk
            references activity
            on update cascade on delete cascade,
    rank        numeric               not null,
    shown       boolean default false not null
);