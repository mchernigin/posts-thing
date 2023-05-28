create table posts (
    id                serial    primary key,
    title             text      not null,
    content           text      not null,
    publish_timestamp timestamp not null
);

create table authors (
    id      serial   primary key,
    name    text     not null,
    surname text     not null,
    website text     null,
    status  text     null
);
