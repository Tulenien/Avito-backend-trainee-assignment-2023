create database avitodb;

\c avitodb

create table users (
    id serial
);

create table segments (
    id serial,
    name text
);

create table users_segments (
    user_id integer,
    segment_id integer,
    constraint fk_users
        foreign_key(user_id)
        references users(id),
    constraint fk_segments
        foreign_key(segment_id)
        references segmetns(id)
);