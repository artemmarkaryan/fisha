create table city
(
    name text,
    lat  numeric,
    lon  numeric,
    id   serial
        constraint city_pk primary key
);