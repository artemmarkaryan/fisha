create table if not exists "activity_date"
(
    id          serial
        constraint activity_date__id_pk primary key,

    activity_id int not null,
    constraint activity_date__activity_id_fk foreign key (activity_id) references "activity" (id) on delete cascade on update cascade,

    start_date  date,
    end_date    date
)