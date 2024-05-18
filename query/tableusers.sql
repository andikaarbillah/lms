create table users(
    id varchar(255) primary key,
    username varchar(55) not null,
    email varchar(55) not null unique,
    password varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp not null,
    is_deleted boolean default false
);