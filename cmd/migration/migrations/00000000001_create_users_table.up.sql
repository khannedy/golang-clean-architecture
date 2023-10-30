create table users
(
    id         varchar(100) not null,
    name       varchar(100) not null,
    password   varchar(100) not null,
    token      varchar(100) null,
    created_at bigint       not null default current_timestamp,
    updated_at bigint       not null default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB;