CREATE TYPE gender AS ENUM ('male', 'female', 'other');

CREATE TABLE IF NOT EXISTS clients (
     id uuid primary key ,
     name varchar(50) not null ,
     surname varchar(50) not null ,
     patronymic varchar(50) ,
     Gender gender not null ,
     age int not null,
     country_id varchar(25) not null
);