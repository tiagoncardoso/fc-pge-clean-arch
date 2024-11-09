CREATE DATABASE IF NOT EXISTS fc_challenge;

USE fc_challenge;

CREATE TABLE IF NOT EXISTS orders
(
    id          varchar(60) not null
        primary key,
    price       float       not null,
    tax         float       not null,
    final_price float       not null
);

