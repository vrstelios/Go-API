create table tourism(
    id         uuid         not null primary key,
    tourism_gr double precision,
    tourism    double precision,
    area_name  varchar(255) not null unique
);

create table unemployments(
    id           uuid             not null  primary key,
    unemployment double precision not null,
    area_name    varchar(255)     not null unique
);

create table deposits(
    id                     uuid         not null primary key,
    deposit_in_one_million integer      not null,
    city_name              varchar(255) not null unique
);

create table city(
    id         uuid          not null primary key,
    name       varchar(255)  not null,
    latitude   numeric(9, 6) not null,
    longitude  numeric(9, 6) not null,
    capital    varchar(255)  not null,
    population integer,
    area_name  varchar(255)  not null,
    geom       geography(Point, 4326)
);

create table area(
    id          uuid         not null primary key,
    name        varchar(255) not null unique,
    area_length double precision
);

create table populations(
    id         uuid         not null primary key,
    population integer      not null,
    area_name  varchar(255) not null   unique
);

create table airports(
    id        uuid         not null primary key,
    area_name varchar(255) not null,
    code      varchar(5),
    name      varchar(255),
    type      varchar(255),
    geom      geography(Point, 4326)
);