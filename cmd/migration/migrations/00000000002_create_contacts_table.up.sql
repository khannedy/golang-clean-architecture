create table contacts
(
    id         varchar(100) not null,
    first_name varchar(100) not null,
    last_name  varchar(100) null,
    email      varchar(100) null,
    phone      varchar(100) null,
    user_id    varchar(100) not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key fk_contacts_user_id (user_id) references users (id)
) engine = innodb;