create table users
(
    id         varchar(100)                        not null,
    username   varchar(100)                        not null,
    name       varchar(100)                        not null,
    password   varchar(100)                        not null,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp on update current_timestamp not null,
    constraint username_unique unique (username),
    primary key (id)
) engine = InnoDB;