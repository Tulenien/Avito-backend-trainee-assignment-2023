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
    segment_id integer
);