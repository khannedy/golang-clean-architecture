create table addresses
(
    id          varchar(100) not null,
    contact_id  varchar(100) not null,
    street      varchar(255),
    city        varchar(255),
    province    varchar(255),
    postal_code varchar(10),
    country     varchar(100),
    created_at  bigint       not null,
    updated_at  bigint       not null,
    primary key (id),
    foreign key fk_addresses_contact_id (contact_id) references contacts (id)
) engine = innodb;