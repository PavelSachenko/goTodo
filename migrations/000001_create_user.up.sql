create table users
(
    id            int auto_increment,
    name          varchar(255) not null,
    password_hash varchar(150) not null,
    constraint users_pk
        primary key (id)
);

create unique index users_id_uindex
    on users (id);

create unique index users_name_uindex
    on users (name);

